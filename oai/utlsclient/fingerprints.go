package utlsclient

import (
	"errors"

	utls "github.com/refraction-networking/utls"
)

// parseJA3 maps a string ID to a uTLS fingerprint
func parseJA3(name string) (utls.ClientHelloID, error) {
	switch name {
	case "chrome":
		return utls.HelloChrome_Auto, nil
	case "firefox":
		return utls.HelloFirefox_Auto, nil
	case "ios":
		return utls.HelloIOS_Auto, nil
	case "random":
		return utls.HelloRandomized, nil
	default:
		return utls.ClientHelloID{}, errors.New("unsupported JA3/fingerprint name")
	}
}
