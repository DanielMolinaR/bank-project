package server

import (
	"encoding/json"
	"net/http"
)

type TransferRequest struct {
	ToAccount int `json:"to_account"`
	Amount    int `json:"amount"`
}

func (s *APIServer) HandleTransfer(w http.ResponseWriter, r *http.Request) error {
	transferRes := &TransferRequest{}
	if err := json.NewDecoder(r.Body).Decode(transferRes); err != nil {
		return err
	}
	defer r.Body.Close()

	return WriteJSON(w, http.StatusOK, transferRes)
}
