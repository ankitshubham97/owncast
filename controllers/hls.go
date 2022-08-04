package controllers

import (
	"fmt"
	"net/http"
	"path"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"

	"github.com/owncast/owncast/config"
	"github.com/owncast/owncast/core"
	"github.com/owncast/owncast/core/data"
	"github.com/owncast/owncast/models"
	"github.com/owncast/owncast/router/middleware"
	"github.com/owncast/owncast/utils"
)

type AuthJwtClaims struct {
	NftContractAddress string `json:"nftContractAddress"`
	NftId string `json:"nftId"`
	WalletPublicAddress string `json:"walletPublicAddress"`
	jwt.StandardClaims
}

// HandleHLSRequest will manage all requests to HLS content.
func HandleHLSRequest(w http.ResponseWriter, r *http.Request) {
	// w.WriteHeader(http.StatusNotFound)
	// return
	authCookie, err := r.Cookie("Authorization")
	if err != nil {
		log.Errorf("Error occured while reading cookie")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Infof("Cookie: %s", authCookie.Value)
	token, err := jwt.ParseWithClaims(authCookie.Value, &AuthJwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		// return hmacSampleSecret, nil
		return []byte("jwt-secret"), nil
	})

	if claims, ok := token.Claims.(*AuthJwtClaims); ok && token.Valid {
		fmt.Println(claims.NftId, claims.NftContractAddress)
		if (claims.NftContractAddress != "0x8437ee943b49945a7270109277942defe30fac25") || (claims.NftId != "0") {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		fmt.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Sanity check to limit requests to HLS file types.
	if filepath.Ext(r.URL.Path) != ".m3u8" && filepath.Ext(r.URL.Path) != ".ts" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	requestedPath := r.URL.Path
	relativePath := strings.Replace(requestedPath, "/hls/", "", 1)
	fullPath := filepath.Join(config.HLSStoragePath, relativePath)

	// If using external storage then only allow requests for the
	// master playlist at stream.m3u8, no variants or segments.
	if data.GetS3Config().Enabled && relativePath != "stream.m3u8" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Handle playlists
	if path.Ext(r.URL.Path) == ".m3u8" {
		// Playlists should never be cached.
		middleware.DisableCache(w)

		// Force the correct content type
		w.Header().Set("Content-Type", "application/x-mpegURL")

		// Use this as an opportunity to mark this viewer as active.
		viewer := models.GenerateViewerFromRequest(r)
		core.SetViewerActive(&viewer)
	} else {
		cacheTime := utils.GetCacheDurationSecondsForPath(relativePath)
		w.Header().Set("Cache-Control", "public, max-age="+strconv.Itoa(cacheTime))
	}

	middleware.EnableCors(w)
	http.ServeFile(w, r, fullPath)
}
