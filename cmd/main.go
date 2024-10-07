package main

import (
	"log"

	"github.com/skuril-bobishku/song-library/internal/server"
	env "github.com/skuril-bobishku/song-library/pkg/systems"
)

func main() {
	port := env.GetEnvString("SERVER_PORT")
	log.Printf("Запуск сервера на порт %s", port)
	server.StartHTTP(port)
}
