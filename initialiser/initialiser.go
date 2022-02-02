package initialiser

import (
	"net/http"

	"github.com/ONSdigital/books-api/api"
	"github.com/ONSdigital/books-api/interfaces"
	dpHttp "github.com/ONSdigital/dp-net/http"
	"github.com/gorilla/mux"
)

type Service struct {
	Server interfaces.HTTPServer
	Router *mux.Router
	API    *api.API
}

func GetHTTPServer(bindAddr string, router http.Handler) interfaces.HTTPServer {
	httpServer := dpHttp.NewServer(bindAddr, router)
	return httpServer
}
