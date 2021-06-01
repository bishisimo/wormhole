// @time : 2021/6/1 14:28
// @author: bishisimo
// @describe:
package task

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"wormhole/protos/redux"
)

type Client struct {
	*grpc.ClientConn
	redux.ReduxClient
}

func (self Client) Close() error {
	return self.ClientConn.Close()
}

func NewLocalClient() *Client {
	port := viper.GetString("self_port")
	conn, err := grpc.Dial("localhost:"+port, grpc.WithInsecure())
	if err != nil {
		return nil
	}
	c := redux.NewReduxClient(conn)
	return &Client{
		ClientConn:  conn,
		ReduxClient: c,
	}
}
