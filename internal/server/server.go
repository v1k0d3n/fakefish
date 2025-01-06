package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/v1k0d3n/fakefish/pkg/api"
	"github.com/v1k0d3n/fakefish/pkg/config"
)

type Server struct {
	Config *config.Config
}

func NewServer(cfg *config.Config) *Server {
	return &Server{Config: cfg}
}

func (s *Server) Start() {
	api.InitHandlers(s.Config)

	addr := fmt.Sprintf(":%d", s.Config.Server.Port)
	log.Printf("Starting server on %s", addr)

	if s.Config.TLS.Enabled {
		log.Fatal(http.ListenAndServeTLS(addr, s.Config.TLS.ClientCert, s.Config.TLS.ClientKey, s.routes()))
	} else {
		log.Fatal(http.ListenAndServe(addr, s.routes()))
	}
}

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/redfish/v1/", api.RootHandler)

	mux.HandleFunc("/redfish/v1/Systems", api.SystemsHandler)
	mux.HandleFunc("/redfish/v1/Systems/1", api.SystemsHandler)
	mux.HandleFunc("/redfish/v1/Managers", api.ManagersHandler)
	mux.HandleFunc("/redfish/v1/Managers/1", api.ManagersHandler)
	mux.HandleFunc("/redfish/v1/Managers/1/VirtualMedia/CD/Actions/VirtualMedia.InsertMedia", api.MountCdHandler)
	mux.HandleFunc("/redfish/v1/Managers/1/Actions/Manager.Reset", api.PowerOffHandler)
	mux.HandleFunc("/redfish/v1/Managers/1/Actions/Manager.Start", api.PowerOnHandler)
	mux.HandleFunc("/redfish/v1/Managers/1/VirtualMedia/CD/Actions/VirtualMedia.EjectMedia", api.UnmountCdHandler)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	return mux
}
