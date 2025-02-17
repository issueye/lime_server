package websocket

import (
	"fmt"
	"lime/internal/global"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

const (
	// 心跳间隔时间，从30秒改为5分钟
	heartbeatInterval = 5 * time.Minute
	// 写入超时时间，从10秒改为2分钟
	writeTimeout = 2 * time.Minute
	// 读取超时时间，从1分钟改为10分钟
	readTimeout = 10 * time.Minute
)

// 单例模式
var (
	once      sync.Once
	wsManager *WebSocketServer
)

// GetWebSocketManager 获取WebSocketServer单例对象
func GetWebSocketManager() *WebSocketServer {
	once.Do(func() {
		wsManager = NewWebSocketManager()
	})
	return wsManager
}

// WebSocketServer 结构体用于管理WebSocket服务端整体相关功能，包括客户端连接管理
type WebSocketServer struct {
	upgrader   websocket.Upgrader
	clients    map[string]*WebSocketClient
	clientsMu  sync.RWMutex // 使用读写锁提升性能
	shutdownCh chan struct{}
	isShutdown bool
	shutdownMu sync.RWMutex
}

// NewWebSocketManager 创建WebSocketServer实例并初始化配置
func NewWebSocketManager() *WebSocketServer {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			// TODO: 在生产环境中应该实现更严格的检查
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws := &WebSocketServer{
		upgrader:   upgrader,
		clients:    make(map[string]*WebSocketClient),
		clientsMu:  sync.RWMutex{},
		shutdownCh: make(chan struct{}),
		isShutdown: false,
	}

	// 启动定期清理不活跃连接的goroutine
	go ws.periodicCleanup()

	return ws
}

// AddClient 将客户端添加到管理列表中
func (s *WebSocketServer) AddClient(client *WebSocketClient) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	client.ConnID = fmt.Sprintf("%s_%s", client.ClientID, client.TaskID)
	s.clients[client.ConnID] = client

	global.Logger.Info("新客户端连接",
		zap.String("client_id", client.ClientID),
		zap.String("conn_id", client.ConnID))
}

// RemoveClient 从管理列表中移除客户端
func (s *WebSocketServer) RemoveClient(client *WebSocketClient) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	if _, exists := s.clients[client.ConnID]; exists {
		delete(s.clients, client.ConnID)
		global.Logger.Info("客户端断开连接",
			zap.String("client_id", client.ClientID),
			zap.String("conn_id", client.ConnID))
	}
}

// Broadcast 向所有连接的客户端广播消息
func (s *WebSocketServer) Broadcast(message []byte, messageType int) {
	s.clientsMu.RLock()
	defer s.clientsMu.RUnlock()

	for _, client := range s.clients {
		err := s.sendMessageToClient(client, message, messageType)
		if err != nil {
			global.Logger.Error("广播消息失败",
				zap.String("client_id", client.ClientID),
				zap.Error(err))
		}
	}
}

// GetClientCount 获取当前连接的客户端数量
func (s *WebSocketServer) GetClientCount() int {
	s.clientsMu.RLock()
	defer s.clientsMu.RUnlock()
	return len(s.clients)
}

// GetClientByID 根据客户端ID查找客户端实例
func (s *WebSocketServer) GetClientByID(connID string) *WebSocketClient {
	s.clientsMu.RLock()
	defer s.clientsMu.RUnlock()
	return s.clients[connID]
}

// CleanInactiveClients 清理长时间不活跃的客户端
func (s *WebSocketServer) CleanInactiveClients(timeout time.Duration) {
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	currentTime := time.Now().Unix()
	for connID, client := range s.clients {
		if currentTime-client.GetLastActiveTime() > int64(timeout.Seconds()) {
			delete(s.clients, connID)
			client.Close()
			global.Logger.Info("清理不活跃客户端",
				zap.String("client_id", client.ClientID),
				zap.String("conn_id", client.ConnID))
		}
	}
}

// SendMessage 向指定客户端发送消息
func (s *WebSocketServer) SendMessage(connID string, message []byte, messageType int) error {
	client := s.GetClientByID(connID)
	if client == nil {
		return fmt.Errorf("client not found: %s", connID)
	}

	return s.sendMessageToClient(client, message, messageType)
}

// sendMessageToClient 向客户端发送消息的内部方法
func (s *WebSocketServer) sendMessageToClient(client *WebSocketClient, message []byte, messageType int) error {
	if client.conn == nil {
		return fmt.Errorf("connection is nil for client: %s", client.ConnID)
	}

	client.conn.SetWriteDeadline(time.Now().Add(writeTimeout))
	err := client.conn.WriteMessage(messageType, message)
	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	return nil
}

// periodicCleanup 定期清理不活跃的连接
func (s *WebSocketServer) periodicCleanup() {
	// 清理检查间隔从5分钟改为15分钟
	ticker := time.NewTicker(15 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// 不活跃超时时间从10分钟改为30分钟
			s.CleanInactiveClients(30 * time.Minute)
		case <-s.shutdownCh:
			return
		}
	}
}

// Shutdown 优雅关闭服务
func (s *WebSocketServer) Shutdown() {
	s.shutdownMu.Lock()
	defer s.shutdownMu.Unlock()

	if s.isShutdown {
		return
	}

	s.isShutdown = true
	close(s.shutdownCh)

	// 关闭所有客户端连接
	s.clientsMu.Lock()
	defer s.clientsMu.Unlock()

	for _, client := range s.clients {
		client.Close()
	}

	global.Logger.Info("WebSocket服务已关闭")
}

// IsShutdown 检查服务是否已关闭
func (s *WebSocketServer) IsShutdown() bool {
	s.shutdownMu.RLock()
	defer s.shutdownMu.RUnlock()
	return s.isShutdown
}
