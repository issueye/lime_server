package websocket

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"lime/internal/global"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

// MessageType 消息类型
type MessageType int

const (
	TextMessage MessageType = iota + 1
	JsonMessage
	BinaryMessage
)

// Message 定义WebSocket消息结构
type Message struct {
	Type    MessageType `json:"type"`
	Content interface{} `json:"content"`
	Time    time.Time   `json:"time"`
}

// HandleWebSocket 处理WebSocket连接的方法
func (s *WebSocketServer) HandleWebSocket(c *gin.Context) {
	// 验证必要参数
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少[id]参数"})
		return
	}

	pId := c.Query("project_id")
	if pId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "缺少[project_id]参数"})
		return
	}

	// 验证客户端是否已存在
	connID := fmt.Sprintf("%s_%s", id, pId)
	if existingClient := s.GetClientByID(connID); existingClient != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "客户端连接已存在"})
		return
	}

	// 升级HTTP连接为WebSocket
	conn, err := s.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		global.Logger.Error("升级WebSocket连接失败",
			zap.Error(err),
			zap.String("id", id),
			zap.String("project_id", pId))
		return
	}

	// 设置读写超时
	conn.SetReadDeadline(time.Now().Add(readTimeout))
	conn.SetWriteDeadline(time.Now().Add(writeTimeout))

	// 创建新的客户端实例
	client := NewWebSocketClient(conn, fmt.Sprintf("client_%d", time.Now().UnixNano()), id, pId)
	s.AddClient(client)

	// 资源清理
	defer func() {
		s.RemoveClient(client)
		client.Close()
	}()

	// 发送欢迎消息
	welcomeMsg := Message{
		Type:    JsonMessage,
		Content: gin.H{"message": "欢迎使用WebSocket服务！", "id": id},
		Time:    time.Now(),
	}

	if err := s.sendJsonMessage(client, welcomeMsg); err != nil {
		global.Logger.Error("发送欢迎消息失败",
			zap.Error(err),
			zap.String("id", id))
		return
	}

	// 消息处理循环
	for {
		messageType, data, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				global.Logger.Error("读取消息错误",
					zap.Error(err),
					zap.String("id", id))
			}
			break
		}

		// 更新客户端最后活动时间
		client.UpdateLastActive()

		// 处理不同类型的消息
		switch messageType {
		case websocket.TextMessage:
			if err := s.handleTextMessage(client, data); err != nil {
				global.Logger.Error("处理文本消息失败",
					zap.Error(err),
					zap.String("id", id))
			}
		case websocket.BinaryMessage:
			if err := s.handleBinaryMessage(client, data); err != nil {
				global.Logger.Error("处理二进制消息失败",
					zap.Error(err),
					zap.String("id", id))
			}
		}
	}
}

// handleTextMessage 处理文本消息
func (s *WebSocketServer) handleTextMessage(client *WebSocketClient, data []byte) error {
	// 尝试解析为JSON消息
	var msg Message
	if err := json.Unmarshal(data, &msg); err == nil {
		// 是JSON消息，进行相应处理
		return s.handleJsonMessage(client, msg)
	}

	// 普通文本消息处理
	response := Message{
		Type:    TextMessage,
		Content: string(data),
		Time:    time.Now(),
	}

	return s.sendJsonMessage(client, response)
}

// handleBinaryMessage 处理二进制消息
func (s *WebSocketServer) handleBinaryMessage(client *WebSocketClient, data []byte) error {
	response := Message{
		Type:    BinaryMessage,
		Content: fmt.Sprintf("收到二进制消息，长度: %d bytes", len(data)),
		Time:    time.Now(),
	}

	return s.sendJsonMessage(client, response)
}

// handleJsonMessage 处理JSON消息
func (s *WebSocketServer) handleJsonMessage(client *WebSocketClient, msg Message) error {
	// 根据消息类型进行处理
	switch msg.Type {
	case TextMessage:
		// 处理文本类型的JSON消息
		response := Message{
			Type: JsonMessage,
			Content: gin.H{
				"message":  "已收到文本消息",
				"original": msg.Content,
			},
			Time: time.Now(),
		}
		return s.sendJsonMessage(client, response)

	case JsonMessage:
		// 处理JSON类型的消息
		response := Message{
			Type: JsonMessage,
			Content: gin.H{
				"message":  "已收到JSON消息",
				"original": msg.Content,
			},
			Time: time.Now(),
		}
		return s.sendJsonMessage(client, response)

	default:
		return fmt.Errorf("不支持的消息类型: %d", msg.Type)
	}
}

// sendJsonMessage 发送JSON消息
func (s *WebSocketServer) sendJsonMessage(client *WebSocketClient, msg Message) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("消息序列化失败: %v", err)
	}

	return s.sendMessageToClient(client, data, websocket.TextMessage)
}
