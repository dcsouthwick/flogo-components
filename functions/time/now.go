package time

import (
	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
	"time"
)

func init() {
	_ = function.Register(&fnnow{})
}

type fnnow struct {
}

func (fnnow) Name() string {
	return "now"
}

func (fnnow) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeInt64}, true
}

func (fnnow) Eval(params ...interface{}) (interface{}, error) {
	return time.Now().Unix(), nil
}
