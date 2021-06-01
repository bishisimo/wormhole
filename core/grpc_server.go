// @time : 2020/10/13 上午10:12
// @author : bishisimo
// @describe: 业务服务核心
package core

import (
	"context"
	"sync"
	"wormhole/protos/redux"
)

var IsBaked = false

type reduxServer struct {
	redux.UnimplementedReduxServer
	CurrentIndex int
	IndexMap     map[int]*redux.Device
	HostMap      map[string]*redux.Device
	NameMap      map[string]*redux.Device
	EventChan    chan *redux.Heat
	Ctx          context.Context
	Cancel       context.CancelFunc
	MessageChan  chan string
	IsBacked     bool
}

var server *reduxServer
var once sync.Once
var SingleChan chan int

func init() {
	SingleChan = make(chan int, 1)
}

func GetReduxServer() *reduxServer {
	once.Do(func() {
		server = &reduxServer{
			CurrentIndex: 0,
			IndexMap:     make(map[int]*redux.Device),
			HostMap:      make(map[string]*redux.Device),
			NameMap:      make(map[string]*redux.Device),
			EventChan:    make(chan *redux.Heat, 255),
			MessageChan:  make(chan string, 255),
			IsBacked:     IsBaked,
		}
		go server.EventHandle()
	})
	return server
}
