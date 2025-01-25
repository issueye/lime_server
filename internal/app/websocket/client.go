package websocket

import (
	"fmt"
	"sync"
	"time"

	"lime/internal/global"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// WebSocketClient 结构体表示一个WebSocket客户端连接
type WebSocketClient struct {
	conn        *websocket.Conn
	Id          string            // 唯一ID
	ClientID    string            // 添加客户端ID
	TaskID      string            // 添加任务ID
	clientMeta  map[string]string // 可以存放客户端的其他元数据，比如用户名等信息
	connectTime int64             // 记录客户端连接时间，方便后续统计等操作
	Status      int               // 状态 0 未连接 1 已连接
	ConnID      string            // 连接ID    ClientID_TaskID
	lastActive  int64             // 最后活跃时间
	closeChan   chan struct{}     // 用于关闭信号
	closeOnce   sync.Once         // 确保只关闭一次
	mu          sync.RWMutex      // 用于保护并发访问
}

// NewWebSocketClient 创建WebSocketClient实例的函数
func NewWebSocketClient(conn *websocket.Conn, id string, clientID string, taskID string) *WebSocketClient {
	now := time.Now().Unix()
	connID := fmt.Sprintf("%s_%s", clientID, taskID)

	client := &WebSocketClient{
		conn:        conn,
		ClientID:    clientID,
		TaskID:      taskID,
		Id:          id,
		ConnID:      connID,
		clientMeta:  make(map[string]string),
		connectTime: now,
		lastActive:  now,
		closeChan:   make(chan struct{}),
		Status:      1, // 设置为已连接状态
	}

	// 启动心跳检测
	go client.startHeartbeat()

	return client
}

// UpdateLastActive 更新最后活跃时间
func (c *WebSocketClient) UpdateLastActive() {
	c.mu.Lock()
	c.lastActive = time.Now().Unix()
	c.mu.Unlock()
}

// GetLastActiveTime 获取最后活跃时间
func (c *WebSocketClient) GetLastActiveTime() int64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastActive
}

// Close 关闭客户端连接
func (c *WebSocketClient) Close() {
	c.closeOnce.Do(func() {
		if c.conn != nil {
			global.Logger.Info("关闭客户端连接",
				zap.String("client_id", c.ClientID),
				zap.String("conn_id", c.ConnID))

			c.Status = 0 // 设置为未连接状态
			close(c.closeChan)
			c.conn.Close()
		}
	})
}

// IsConnected 检查客户端是否已连接
func (c *WebSocketClient) IsConnected() bool {
	return c.Status == 1
}

// startHeartbeat 开始心跳检测
func (c *WebSocketClient) startHeartbeat() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, []byte{}, time.Now().Add(time.Second)); err != nil {
				global.Logger.Error("心跳检测失败",
					zap.String("client_id", c.ClientID),
					zap.String("conn_id", c.ConnID),
					zap.Error(err))
				c.Close()
				return
			}
			c.UpdateLastActive()
		case <-c.closeChan:
			return
		}
	}
}

// SetMeta 设置客户端元数据
func (c *WebSocketClient) SetMeta(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.clientMeta[key] = value
}

// GetMeta 获取客户端元数据
func (c *WebSocketClient) GetMeta(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.clientMeta[key]
	return value, exists
}
