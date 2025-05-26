package main

import (
	"github.com/alnovi/holidays/internal/app/server"
)

func main() {
	server.NewApp(nil).Start(nil)
}
