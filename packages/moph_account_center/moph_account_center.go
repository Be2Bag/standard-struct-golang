package moph_account_center

import (
	"standard-struct-golang/packages/requests"
	"sync"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const packageName = "moph_account_center"

type Client struct {
	credential   Credential
	url          string
	username     string
	hashPassword string
	hCode        string
	http         *requests.HttpClient
	log          *logrus.Entry
	tracer       trace.Tracer
}

func New(url string, username string, hashPassword string, hCode string, http *requests.HttpClient, logger *logrus.Entry) *Client {
	return &Client{
		credential: Credential{
			mutex: &sync.Mutex{},
		},
		url:          url,
		username:     username,
		hashPassword: hashPassword,
		hCode:        hCode,
		http:         http,
		log:          logger.Dup().WithField("package", packageName),
		tracer:       otel.Tracer(packageName),
	}
}
