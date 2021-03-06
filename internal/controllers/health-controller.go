package controllers

import (
	"fmt"
	"net/http"
	"snow-auth-service/configs"
)

func HealthController(w http.ResponseWriter, r *http.Request) {
	// Set the return Content-Type as JSON like before
	w.Header().Set("Content-Type", "application/json")

	// Change the response depending on the method being requested
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fmt.Sprintf(`{"status": "ok", "version": "%v"}`, configs.GetVersion())))
		if err != nil {
			panic(err)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, err := w.Write([]byte(`{"message": "This HTTP Method is not allowed on this route!"}`))
		if err != nil {
			panic(err)
		}
	}
}
