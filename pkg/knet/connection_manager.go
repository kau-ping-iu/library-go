package knet

import "sync"

const connectionMapNum = 32

type connectionMap struct {
	sync.Mutex // 使用 Mutex 進行鎖定，減少競爭
	connects   map[uint64]Connection
	disposed   bool      // 資源是否已回收
	closeOnce  sync.Once // 使用 sync.Once 保證 Close 僅呼叫一次
}

// ConnectionManager 連線管理器
type ConnectionManager struct {
	connectionMaps [connectionMapNum]connectionMap
	disposeFlag    bool // 是否在處理釋放資源
}

// Dispose 釋放資源
func (m *ConnectionManager) Dispose() {
	m.disposeFlag = true
	for i := 0; i < connectionMapNum; i++ {
		c := &m.connectionMaps[i]
		c.Lock()
		c.disposed = true // 設定資源為已回收
		for _, conn := range c.connects {
			_ = conn.Close() // 關閉連線
		}
		c.Unlock()
	}
}

// GetConnection 取得連線...
func (m *ConnectionManager) GetConnection(connID uint64) Connection {
	cs := &m.connectionMaps[connID%connectionMapNum]
	cs.Lock()
	defer cs.Unlock()
	conn, _ := cs.connects[connID]
	return conn
}

func (m *ConnectionManager) putConnection(conn Connection) {
	cs := &m.connectionMaps[conn.GetConnID()%connectionMapNum]
	cs.Lock()
	defer cs.Unlock()
	if cs.disposed {
		_ = conn.Close() // 如果資源已回收，關閉連線
		return
	}
	cs.connects[conn.GetConnID()] = conn
}

func (m *ConnectionManager) delConnection(conn Connection) {
	if m.disposeFlag {
		return // 如果正在處理釋放資源，則不進行刪除操作
	}
	cs := &m.connectionMaps[conn.GetConnID()%connectionMapNum]
	cs.Lock()
	defer cs.Unlock()
	delete(cs.connects, conn.GetConnID()) // 刪除連線
}

func NewConnectionManager() *ConnectionManager {
	manager := &ConnectionManager{}
	for i := 0; i < len(manager.connectionMaps); i++ {
		manager.connectionMaps[i].connects = make(map[uint64]Connection) // 初始化連線地圖
	}
	return manager
}
