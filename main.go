package main

import (
	"github.com/omise/omise-go"
	"net/http"
)

const (
	// Read these from environment variables or configuration files!
	OmisePublicKey = "pkey_test_521w1g1t7w4x4rd22z0"
	OmiseSecretKey = "skey_test_521w1g1t6yh7sx4pu8n"
)

func main() {
	
	if OmisePublicKey == "" && OmiseSecretKey == "" {
		panic("Please set OMISE_SKEY")
	}

	client, e := omise.NewClient(OmisePublicKey, OmiseSecretKey)
	if e != nil {
		panic(e)
	}

	if e := http.ListenAndServe(":8080", &TamboonHandler{client}); e != nil {
		panic(e)
	}
}