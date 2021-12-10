package go_doppler

import (
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

// query doppler config via its environment variable and returns its secrets
// variable DOPPLER_TOKEN has to be initialised
func getSecrets() (*resty.Response, error) {
	var token = os.Getenv("DOPPLER_TOKEN")
	getURL := fmt.Sprintf("https://%s@api.doppler.com/v3/configs/config/secrets/download?format=json", token)

	// ----------------------------------------- request
	req := resty.New()

	// to debug use: req.SetDebug(true).R().
	secrets, err := req.R().SetHeader("Content-type", "application/x-www-form-urlencoded").Get(getURL)
	if err != nil {
		return nil, err
	}

	return secrets, nil
}
