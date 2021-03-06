package main

import (
	"fmt"
	"net/http"
	"sy/pkg/setting"
	"sy/pkg/task"
	"sy/routers"
)

func main() {
	task.Listen()
	fmt.Println("服务开始")
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: setting.MaxHeaderBytes,
	}
	s.ListenAndServe()
}
