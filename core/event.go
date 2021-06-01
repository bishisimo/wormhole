// @time : 2021/5/31 15:51
// @author: bishisimo
// @describe: 事件处理模块
package core

import (
	"wormhole/event"
	"wormhole/protos/redux"
)

func (self *reduxServer) EventAdd(heat *redux.Heat) {
	self.EventChan <- heat
}
func (self *reduxServer) EventHandle() {
	for heat := range self.EventChan {
		if heat.Event == event.OfflineEvent {
			self.OfflineDevice(heat)
		} else {
			self.OnlineDevice(heat)
		}
	}
}
