package main

import (
	server "projeto404/src/Api/Server"
)

func main() {

	// Define rotas e roda servidor
	rotas := server.RodarServidor()
	rotas.Run()
}
