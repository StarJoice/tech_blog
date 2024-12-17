package ioc

import (
	"github.com/google/wire"
)

var BaseSet = wire.NewSet(InitDB)
