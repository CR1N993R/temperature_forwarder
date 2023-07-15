package logs

import (
	"io"
	"loki-log-creator/config"
	"loki-log-creator/loki"
	"net/http"
)

type Controller struct {
	Context       string
	Configuration config.Config
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	_ = r.Body.Close()
	err := loki.SendLogToLoki(string(body), c.Context, c.Configuration.LokiInstances)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
