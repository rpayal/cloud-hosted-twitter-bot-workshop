package twitter_auth

import (
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	logr "github.com/sirupsen/logrus"
)

// Credentials struct contains API credentials pulled from env vars:
type Credentials struct {
	ApiKey            string
	ApiSecretKey      string
	AccessToken       string
	AccessTokenSecret string
}

func GetCredentials() Credentials {
	// Populating the struct - value semantic construction
	creds := Credentials{
		ApiKey:            os.Getenv("API_KEY"),
		ApiSecretKey:      os.Getenv("API_SECRET_KEY"),
		AccessToken:       os.Getenv("ACCESS_TOKEN"),
		AccessTokenSecret: os.Getenv("ACCESS_TOKEN_SECRET"),
	}
	return creds
}

/* GetUserClient:
Input = credentials
Return = client, error
*/
func GetUserClient(creds *Credentials) (*twitter.Client, error) {

	// Create a new config & token using the data stored in 'creds'
	config := oauth1.NewConfig(creds.ApiKey, creds.ApiSecretKey)
	token := oauth1.NewToken(creds.AccessToken, creds.AccessTokenSecret)

	// Create a new http client
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	verifyParams := &twitter.AccountVerifyParams{
		SkipStatus:   twitter.Bool(true),
		IncludeEmail: twitter.Bool(true),
	}

	// Verify the user credentials
	user, _, err := client.Accounts.VerifyCredentials(verifyParams)
	if err != nil {
		logr.Error(err)
		return nil, err
	}

	logr.Infof("User Account Info:\n%+v\n", user)
	return client, nil
}
