package hermes

import (
	"net/http"
)

func setupResponse(worker *Worker, w *http.ResponseWriter, r *http.Request) {
	if worker.CORS {
		(*w).Header().Set("Access-Control-Allow-Origin", "*")
		(*w).Header().Set("Access-Control-Allow-ZMethods", "POST, GET, OPTIONS, PUT, DELETE")
		(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}
}
