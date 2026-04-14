package client

import (
	"fmt"
	"go-ent-demo/util/log"
)

func logging(args ...any) {
	log.Debugf("[ORM CORE] %s", fmt.Sprint(args...))
}
