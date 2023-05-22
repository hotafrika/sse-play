package sse

import (
	"log"
	"net/http"
	"os"

	"github.com/alexandrevicenzi/go-sse"
)

type SSEService struct {
	server *sse.Server
}

func NewSSEService() SSEService {
	return SSEService{
		server: sse.NewServer(
			&sse.Options{
				RetryInterval: 100,
				Headers: map[string]string{
					"Access-Control-Allow-Origin": "*",
				},
				Logger: log.New(os.Stdout, "go-sse: ", log.LstdFlags),
			},
		),
	}
}

func (s SSEService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.server.ServeHTTP(w, r)
}
