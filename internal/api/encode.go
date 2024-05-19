package api

import (
	"encoding/json"
	"net/http"
)

func encodeJson(w http.ResponseWriter, payload any) {
	out, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
