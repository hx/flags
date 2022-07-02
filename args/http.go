package args

import (
	"errors"
	"fmt"
	"github.com/hx/flags/app"
	"github.com/hx/flags/interfaces"
	"net/url"
	"strings"
)

func init() {
	registrar["http"] = func(args []string, config *app.Config) (string, error) {
		if len(args) != 1 {
			return "", errors.New("expected exactly 1 argument")
		}
		str := args[0]
		if !strings.Contains(str, "://") {
			str = "http://" + str
		}
		u, err := url.Parse(str)
		if err != nil {
			return "", err
		}
		passPhrase := u.User.String()
		server := interfaces.NewHttpServer(u.Host, passPhrase)
		config.Input(server)
		config.Output(server)
		passPhraseMessage := "no passphrase"
		if passPhrase != "" {
			passPhraseMessage = "passphrase: " + passPhrase
		}
		config.Input(server).Output(server)
		return fmt.Sprintf("Server listening on %s with %s", u.Host, passPhraseMessage), nil
	}
}
