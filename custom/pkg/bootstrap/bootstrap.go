package bootstrap

import (
	"github.com/awitas/myapp/checksum"
	"github.com/viant/xdatly/types/core"
	"reflect"
)

var PackageName = "bootstrap"

func init() {
	core.RegisterType(PackageName, "Bootstrap", reflect.TypeOf(struct{}{}), checksum.GeneratedTime)
}
