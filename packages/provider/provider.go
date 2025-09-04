package provider

import (
	"standard-struct-golang/packages/requests"

	"github.com/sirupsen/logrus"
)

const packageName = "provider"

type Provider struct {
	http     *requests.HttpClient
	url      string
	redirect string
	id       string
	secret   string
	log      *logrus.Entry
	Timeout  int
}

func New(http *requests.HttpClient, url string, redirectUrl string, clientId string, clientSecret string, logger *logrus.Entry, timeout int) *Provider {
	return &Provider{
		http:     http,
		url:      url,
		redirect: redirectUrl,
		id:       clientId,
		secret:   clientSecret,
		log:      logger.Dup().WithField("package", packageName),
		Timeout:  timeout,
	}
}
