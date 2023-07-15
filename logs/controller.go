package logs

import (
	"io"
	"net/http"
	"temperature_forwarder/loki"
)

type Controller struct {
	Context string
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	_ = r.Body.Close()
	err := loki.SendLogToLoki(string(body), c.Context)
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
}
