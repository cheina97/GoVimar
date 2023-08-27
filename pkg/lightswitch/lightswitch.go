package lightswitch

import (
	"fmt"
	"io"
	"net/http"

	"github.com/cheina97/govimar/pkg/config"
)

type Switch struct {
	config *config.Switch
	key    string
}

func NewSwitch(scfg *config.Switch, key string) *Switch {
	return &Switch{
		config: scfg,
		key:    key}
}

func (s *Switch) sendEvent(event config.Event) (status, body string, err error) {
	res, err := http.Post(
		fmt.Sprintf("https://maker.ifttt.com/trigger/%s/with/key/%s", event, s.key),
		"application/json",
		nil,
	)
	if err != nil {
		return "", "", err
	}
	resbytes, err := io.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}
	return res.Status, string(resbytes), nil
}

// On sends a request to the IFTTT Maker channel to turn on the switch.
func (s *Switch) On() (status, body string, err error) {
	return s.sendEvent(s.config.Events.On)
}

// Off sends a request to the IFTTT Maker channel to turn off the switch.
func (s *Switch) Off() (status, body string, err error) {
	return s.sendEvent(s.config.Events.Off)
}
