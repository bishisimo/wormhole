// @time : 2021/6/1 16:16
// @author: bishisimo
// @describe:
package utils

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"runtime"
)

func GetCallerName() string {
	pc, _, _, _ := runtime.Caller(2)
	fn, line := runtime.FuncForPC(pc).FileLine(pc)
	name := runtime.FuncForPC(pc).Name()
	p := fmt.Sprintf("%v:%v->%v", fn, line, name)
	return p
}
func HandleErr(err error) {
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "")).Send()
	}
}
func HandleNilOrErr(data interface{}, err error) {
	if err != nil {
		log.Error().Stack().Err(errors.Wrap(err, "")).Send()
	} else if data == nil {
		log.Error().Stack().Err(errors.New("data is nil")).Send()
	}
}
