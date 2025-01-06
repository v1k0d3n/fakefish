package api

import (
	"encoding/json"
	"net/http"

	"github.com/v1k0d3n/fakefish/pkg/config"
	"github.com/v1k0d3n/fakefish/pkg/scripts"
)

var cfg *config.Config

func InitHandlers(c *config.Config) {
	cfg = c
}

type InsertMediaRequest struct {
	Image string `json:"Image"`
}

func BootFromCdOnceHandler(w http.ResponseWriter, r *http.Request) {
	scriptPath, ok := cfg.Commands["bootFromCdOnce"]
	if !ok {
		http.Error(w, "Boot from CD once script not found", http.StatusInternalServerError)
		return
	}

	err := scripts.ExecuteScript(scriptPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Boot from CD once executed successfully"))
}

func MountCdHandler(w http.ResponseWriter, r *http.Request) {
	var req InsertMediaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	scriptPath, ok := cfg.Commands["mountCd"]
	if !ok {
		http.Error(w, "Mount CD script not found", http.StatusInternalServerError)
		return
	}

	err := scripts.ExecuteScript(scriptPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Mount CD executed successfully"))
}

func PowerOffHandler(w http.ResponseWriter, r *http.Request) {
	scriptPath, ok := cfg.Commands["powerOff"]
	if !ok {
		http.Error(w, "Power off script not found", http.StatusInternalServerError)
		return
	}

	err := scripts.ExecuteScript(scriptPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Power off executed successfully"))
}

func PowerOnHandler(w http.ResponseWriter, r *http.Request) {
	scriptPath, ok := cfg.Commands["powerOn"]
	if !ok {
		http.Error(w, "Power on script not found", http.StatusInternalServerError)
		return
	}

	err := scripts.ExecuteScript(scriptPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Power on executed successfully"))
}

func UnmountCdHandler(w http.ResponseWriter, r *http.Request) {
	scriptPath, ok := cfg.Commands["unmountCd"]
	if !ok {
		http.Error(w, "Unmount CD script not found", http.StatusInternalServerError)
		return
	}

	err := scripts.ExecuteScript(scriptPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Unmount CD executed successfully"))
}
