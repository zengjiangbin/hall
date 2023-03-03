package main

import (
	"github.com/zengjiangbin/driver"
	"github.com/zengjiangbin/hall/core"
)

func main() {
	driver.Drive(core.NewService())
}
