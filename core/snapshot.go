// @time : 2021/5/31 15:52
// @author: bishisimo
// @describe: 服务端快照模块
package core

import (
	"context"
	"github.com/bradleyjkemp/memviz"
	"github.com/sirupsen/logrus"
	"os"
	"os/exec"
	"wormhole/protos/redux"
)

func (self *reduxServer) Snapshot(ctx context.Context, empty *redux.Empty) (*redux.Reply, error) {
	f, err := os.OpenFile("./core.dot", os.O_WRONLY|os.O_CREATE, os.FileMode(0777))
	if err != nil {
		logrus.Fatalln(err.Error())
	}
	memviz.Map(f, self)
	command := exec.Command("dot", "-Tpng", "core.dot", "-o", "core.png")
	err = command.Run()
	if err != nil {
		logrus.Warningln(err.Error())
		return nil, err
	}
	return new(redux.Reply), nil
}
