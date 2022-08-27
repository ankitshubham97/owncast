package controllers

import (
	"net/http"

	"github.com/owncast/owncast/core"
	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/router/middleware"
)

// Ping is fired by a client to show they are still an active viewer.
func Ping(w http.ResponseWriter, r *http.Request) {
	responseCode := middleware.AuthByNft(w, r)
	if responseCode != http.StatusOK {
		w.Header().Set("Nft-Auth", "false")
	} else {
		w.Header().Set("Nft-Auth", "true")
	}
	viewer := models.GenerateViewerFromRequest(r)
	core.SetViewerActive(&viewer)
	w.WriteHeader(responseCode)
}
