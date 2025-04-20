package main

import (
	"diplom/wallet-service/internal/app"
	"diplom/wallet-service/pkg/configs"
	"log"
)

func main() {
	conf, err := configs.LoadConfigs()
	if err != nil {
		log.Fatalf("load conf failed, err:%v", err)
	}

	err = app.Start(conf)
	if err != nil {
		log.Fatalf("start app failed, err:%v", err)
	}
}
