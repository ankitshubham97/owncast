package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
	log "github.com/sirupsen/logrus"
)

type AuthJwtClaims struct {
	NftContractAddress string `json:"nftContractAddress"`
	NftId string `json:"nftId"`
	WalletPublicAddress string `json:"walletPublicAddress"`
	jwt.StandardClaims
}

// EnableCors enables the CORS header on the responses.
func AuthByNft(w http.ResponseWriter, r *http.Request) int {
	authCookie, err := r.Cookie("Authorization")
	if err != nil {
		log.Errorf("Error occured while reading cookie")
		return http.StatusNotFound
	
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
			return http.StatusBadRequest
		
		} else {
			return http.StatusOK
		}
	} else {
		fmt.Println(err)
		return http.StatusBadRequest
	
	}
}
