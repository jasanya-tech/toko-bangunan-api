package main

import (
	customLog "toko-bangunan/internal/logger"
	"toko-bangunan/internal/protocols/http"
)

func main() {
	customLog.InitLogger()
	var http http.HttpImpl
	http.Listen()
}
