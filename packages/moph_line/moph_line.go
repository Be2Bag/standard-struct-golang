package moph_line

import (
	"standard-struct-golang/packages/requests"

	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const packageName = "moph_line"

type Client struct {
	url    string
	http   *requests.HttpClient
	log    *logrus.Entry
	tracer trace.Tracer
}

func New(url string, http *requests.HttpClient, logger *logrus.Entry) *Client {
	return &Client{
		url:    url,
		http:   http,
		log:    logger.Dup().WithField("package", packageName),
		tracer: otel.Tracer(packageName),
	}
}
