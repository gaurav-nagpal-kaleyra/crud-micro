package health

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type resp struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	var result resp

	if isServerHealthy() {
		w.WriteHeader(http.StatusOK)
		result = resp{Status: "ok"}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		result = resp{Status: "Error"}
	}

	err := json.NewEncoder(w).Encode(result)

	if err != nil {
		zap.L().Error("Unable to encode response", zap.Error(err))
	}
}

func isServerHealthy() bool {
	//currently we have no db connections, so returning true
	return true
}
