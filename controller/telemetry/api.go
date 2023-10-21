package telemetry

import (
	"errors"
	"net/http"
	"time"

	"github.com/reef-pi/reef-pi/controller/storage"
	"github.com/reef-pi/reef-pi/controller/utils"
)

// TODO: translate these
// TODO: this is a bit of a hack in that it means someone can't have their
//
//	actual token/password be the string "<stored>", but that seems rare
//	enough a reasonable trade-off for a quick fix to unsaveable form bugs
const PasswordStoredPlaceholder = "<stored>"
const AdafruitIOTokenStoredPlaceholder = "<stored>"

func (t *telemetry) GetConfig(w http.ResponseWriter, req *http.Request) {
	fn := func(_ string) (interface{}, error) {
		var c TelemetryConfig
		if err := t.store.Get(t.bucket, DBKey, &c); err != nil {
			return nil, err
		}
		if c.AdafruitIO.Token != "" {
			c.AdafruitIO.Token = AdafruitIOTokenStoredPlaceholder
		}
		if c.Mailer.Password != "" {
			c.Mailer.Password = PasswordStoredPlaceholder
		}
		return &c, nil
	}
	utils.JSONGetResponse(fn, w, req)
}

func (t *telemetry) UpdateConfig(w http.ResponseWriter, req *http.Request) {
	var c TelemetryConfig

	var existingConfig TelemetryConfig
	var readErr = t.store.Get(t.bucket, DBKey, &existingConfig)
	if readErr != nil {
		if errors.Is(readErr, storage.ErrDoesNotExist) {
			utils.ErrorResponse(http.StatusInternalServerError, "Failed to update. Error: "+readErr.Error(), w)
			return
		}
	}

	fn := func(_ string) error {
		if readErr != nil {
			if c.AdafruitIO.Token == AdafruitIOTokenStoredPlaceholder {
				c.AdafruitIO.Token = existingConfig.AdafruitIO.Token
			}
			if c.Mailer.Password == PasswordStoredPlaceholder {
				c.Mailer.Password = existingConfig.Mailer.Password
			}
		}
		return t.store.Update(t.bucket, DBKey, c)
	}
	utils.JSONUpdateResponse(&c, fn, w, req)
}

func (t *telemetry) SendTestMessage(w http.ResponseWriter, req *http.Request) {
	fn := func(_ string) error {
		_, err := t.Alert("Test email", "This is a test email, generated by reef-pi at: "+time.Now().Format(time.RFC822))
		return err
	}
	utils.JSONDeleteResponse(fn, w, req)
}
