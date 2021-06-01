// @time : 2021/5/31 15:57
// @author: bishisimo
// @describe: 消息扩展
package core

import "wormhole/protos/redux"

//提供本地消息流，给外部程序使用
func (self *reduxServer) GetMessageStream(empty *redux.Empty, streamServer redux.Redux_GetMessageStreamServer) error {
	for message := range self.MessageChan {
		err := streamServer.Send(&redux.Message{Data: message})
		if err != nil {
			return err
		}
	}
	return nil
}

//将消息储存到队列
func (self *reduxServer) SaveMessage(message string) {
	if len(self.MessageChan) > 250 {
		<-self.MessageChan //消息队列慢，丢弃最久的消息
	}
	self.MessageChan <- message
}
