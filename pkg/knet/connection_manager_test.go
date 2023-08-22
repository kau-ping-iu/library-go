package knet

//func TestConnectionManager_Dispose(t *testing.T) {
//	type fields struct {
//		connectionMaps [connectionMapNum]connectionMap
//		disposeFlag    bool
//	}
//	tests := []struct {
//		name   string
//		fields fields
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &ConnectionManager{
//				connectionMaps: tt.fields.connectionMaps,
//				disposeFlag:    tt.fields.disposeFlag,
//			}
//			m.Dispose()
//		})
//	}
//}
//
//func TestConnectionManager_GetConnection(t *testing.T) {
//	type fields struct {
//		connectionMaps [connectionMapNum]connectionMap
//		disposeFlag    bool
//	}
//	type args struct {
//		connID uint64
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//		want   Connection
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &ConnectionManager{
//				connectionMaps: tt.fields.connectionMaps,
//				disposeFlag:    tt.fields.disposeFlag,
//			}
//			if got := m.GetConnection(tt.args.connID); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("GetConnection() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestConnectionManager_delConnection(t *testing.T) {
//	type fields struct {
//		connectionMaps [connectionMapNum]connectionMap
//		disposeFlag    bool
//	}
//	type args struct {
//		conn Connection
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &ConnectionManager{
//				connectionMaps: tt.fields.connectionMaps,
//				disposeFlag:    tt.fields.disposeFlag,
//			}
//			m.delConnection(tt.args.conn)
//		})
//	}
//}
//
//func TestConnectionManager_putConnection(t *testing.T) {
//	type fields struct {
//		connectionMaps [connectionMapNum]connectionMap
//		disposeFlag    bool
//	}
//	type args struct {
//		conn Connection
//	}
//	tests := []struct {
//		name   string
//		fields fields
//		args   args
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			m := &ConnectionManager{
//				connectionMaps: tt.fields.connectionMaps,
//				disposeFlag:    tt.fields.disposeFlag,
//			}
//			m.putConnection(tt.args.conn)
//		})
//	}
//}
//
//func TestNewConnectionManager(t *testing.T) {
//	tests := []struct {
//		name string
//		want *ConnectionManager
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := NewConnectionManager(); !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("NewConnectionManager() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
