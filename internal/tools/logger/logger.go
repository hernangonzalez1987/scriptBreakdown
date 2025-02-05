package httplogger

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

type httpLogger struct {
	logger    *zerolog.Logger
	transport http.RoundTripper
}

func New(logger *zerolog.Logger, transport http.RoundTripper) http.RoundTripper {
	return &httpLogger{logger: logger, transport: transport}
}

func (ref *httpLogger) RoundTrip(req *http.Request) (*http.Response, error) {
	ref.logger.Info().Any("req", req).Msg("http request sent")

	res, err := ref.transport.RoundTrip(req)

	if err != nil {
		ref.logger.Error().Any("res", res).Err(err).Msg("http error response received")
	} else {
		ref.logger.Info().Any("res", res).Msg("http success response received")
	}

	return res, errors.WithStack(err)
}
