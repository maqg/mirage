package main

import (
	"fmt"
	"net/http"
	"octlink/mirage/src/api"
	"octlink/mirage/src/utils"
	"octlink/mirage/src/utils/octlog"
)

func initDebugConfig() {
	octlog.InitDebugConfig(octlog.DEBUG_LEVEL)
}

func initLogConfig() {
	api.InitApiLog(octlog.DEBUG_LEVEL)
}

const (
	HTTP_SERVER = "0.0.0.0"
	HTTP_PORT   = 9999
)

func main() {

	fmt.Println(utils.Version())

	initDebugConfig()
	initLogConfig()

	api := &api.Api{
		Name: "Mirage API Server",
	}

	server := &http.Server{
		Addr:           fmt.Sprintf("%s:%d", HTTP_SERVER, HTTP_PORT),
		Handler:        api.ApiRouter(),
		MaxHeaderBytes: 1 << 20,
	}

	octlog.Warn("Mirage Engine Started\n")

	err := server.ListenAndServe()
	if err != nil {
		octlog.Error("error to listen\n")
	}
}