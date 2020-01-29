package time

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

func init() {
	_ = function.Register(&fnNow{})
}

type fnNow struct {
}

func (fnNow) Name() string {
	return "now"
}

func (fnNow) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt64}, true
}

func (fnNow) Eval(params ...interface{}) (interface{}, error) {
	return time.Now().Unix(), nil
}
