package bootstrap

import (
	"github.com/EdgardoArriagada/hexagonal-go/internal/platform/server"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host = "localhost"
	port = 8080
)

func Run() error {
	srv := server.New(host, port)
	return srv.Run()
}
