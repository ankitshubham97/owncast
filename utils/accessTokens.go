package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"

	log "github.com/sirupsen/logrus"
)

const tokenLength = 32
type response2 struct {
	ErrorCode   string      `json:"errorCode"`
	AccessToken string 	`json:"accessToken"`
}

func GenerateRandomAccessToken() (string, error) {
	return generateRandomString(tokenLength)
}

// GenerateAccessToken will generate and return an access token.
func GenerateAccessToken(nonce string, signature string, walletPublicAddress string, nftContractAddress string, nftId string) (string, error) {
	postBody, _ := json.Marshal(map[string]string{
    "nonce": nonce,
    "signature": signature,
    "walletPublicAddress": walletPublicAddress,
    "nftContractAddress": nftContractAddress,
    "nftId": nftId,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post("http://localhost:3000/token", "application/json", responseBody)

	if err != nil {
			log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
			log.Fatalln(err)
	}
	dat := response2{}
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}
	if err := dat.ErrorCode; err != "" {
		return "", errors.New(err)
	}
	return dat.AccessToken, nil
}

// generateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// generateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func generateRandomString(n int) (string, error) {
	b, err := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b), err
}
