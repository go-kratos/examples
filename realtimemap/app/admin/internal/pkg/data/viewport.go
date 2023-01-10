package data

import (
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-realtimemap/api/admin/v1"
)

type ViewportMap map[websocket.SessionID]*v1.Viewport
