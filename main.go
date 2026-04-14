package main

import (
	"go-ent-demo/config"
	_ "go-ent-demo/entcore/runtime"
	"go-ent-demo/util/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.GetConfig()
	log.InitLog()
	log.Infof("appName: %s, env: %s", cfg.App.Name, cfg.App.Env)
}
