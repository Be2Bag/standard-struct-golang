package health_id

import (
	"standard-struct-golang/packages/requests"

	"github.com/sirupsen/logrus"
)

const packageName = "health_id"

type HealthId struct {
	http              *requests.HttpClient
	url               string
	id                string
	secret            string
	redirectUrl       string
	redirectLocalhost string
	mophUrl           string
	mophClientID      string
	mophSecret        string
	log               *logrus.Entry
	// Timeout           time.Duration
	Timeout int
}

func New(http *requests.HttpClient, url, clientId, clientSecret, redirectUrl, redirectLocalhost, mophUrl, mophClientID, mophSecret string, logger *logrus.Entry, timeout int) *HealthId {
	return &HealthId{
		http:              http,
		url:               url,
		id:                clientId,
		secret:            clientSecret,
		redirectUrl:       redirectUrl,
		redirectLocalhost: redirectLocalhost,
		mophUrl:           mophUrl,
		mophClientID:      mophClientID,
		mophSecret:        mophSecret,
		log:               logger.Dup().WithField("package", packageName),
		// Timeout:           time.Duration(timeout) * time.Second,
		Timeout: timeout,
	}
}
