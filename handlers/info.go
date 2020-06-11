package handlers

import (
	"encoding/json"
	"github.com/drag0ns3c/httpg/sys"
	"net/http"
)

func SystemInfoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sys.New())
}
