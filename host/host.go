package host

import (
	"github.com/kevinburke/ssh_config"
	"github.com/trntv/sshed/keychain"
	"strings"
)

type Host struct {
	Key          string
	Hostname     string
	Port         string
	User         string
	IdentityFile string

	Options map[string]string

	// Params holds every key/value line of the host block in file order,
	// including repeated keys (e.g. LocalForward) and inline comments.
	Params []Param
	// Comment is the comment on the "Host" line itself, if any.
	Comment string

	creds *keychain.Record
}

// Param is a single "Key Value # comment" line of a host block.
type Param struct {
	Key     string
	Value   string
	Comment string
}

func CreateFromConfig(h *ssh_config.Host) *Host {
	hh := &Host{
		Options: make(map[string]string),
		Comment: strings.TrimSpace(h.EOLComment),
	}
	for _, node := range h.Nodes {
		switch node.(type) {
		case *ssh_config.KV:
			c := node.(*ssh_config.KV)
			hh.Params = append(hh.Params, Param{
				Key:     c.Key,
				Value:   strings.TrimSpace(c.Value),
				Comment: strings.TrimSpace(c.Comment),
			})
			switch strings.ToLower(c.Key) {
			case "hostname":
				hh.Hostname = c.Value
			case "port":
				hh.Port = c.Value
			case "user":
				hh.User = c.Value
			case "identityfile":
				hh.IdentityFile = c.Value
			default:
				hh.Options[c.Key] = c.Value
			}
		}
	}

	return hh
}

func (h *Host) Password() (password string) {
	h.readKeychain()
	if h.creds != nil {
		password = h.creds.Password
	}
	return password
}

func (h *Host) PrivateKey() (pk string) {
	h.readKeychain()
	if h.creds != nil {
		pk = h.creds.PrivateKey
	}
	return pk
}

func (h *Host) readKeychain() error {
	var err error
	if h.creds != nil {
		return nil
	}
	h.creds, err = keychain.Get(h.Key)
	return err
}
