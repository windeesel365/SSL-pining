package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"serverside/handler"

	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()

	e.GET("/api/hello", handler.SecuredHandler)

	cert, err := tls.LoadX509KeyPair("cert_and_key/server-cert.pem", "cert_and_key/server-key.pem")
	if err != nil {
		log.Fatalf("failed to load key pair: %v", err)
	}

	server := &http.Server{
		Addr: ":1150",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	log.Println("Starting server on https://localhost:1150")
	if err := e.StartServer(server); err != nil {
		log.Fatalf("server failed to start: %v", err)
	}
}
