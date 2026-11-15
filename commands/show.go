package commands

import (
	"fmt"
	"github.com/mgutz/ansi"
	"github.com/pkg/errors"
	"github.com/trntv/sshed/host"
	"github.com/trntv/sshed/ssh"
	"github.com/urfave/cli"
	"strings"
)

func (cmds *Commands) newShowCommand() cli.Command {
	return cli.Command{
		Name:      "show",
		Usage:     "Shows host",
		ArgsUsage: "<key>",
		Action:    cmds.showAction,
		BashComplete: func(c *cli.Context) {
			// This will complete if no args are passed
			if c.NArg() > 0 {
				return
			}
			cmds.completeWithServers()
		},
	}
}

func (cmds *Commands) showAction(c *cli.Context) (err error) {
	var key string

	if c.NArg() == 0 {
		key, err = cmds.askServerKey()
		if err != nil {
			return err
		}
	} else {
		key = c.Args().First()
	}

	srv := ssh.Config.Get(key)
	if srv == nil {
		return errors.New("host not found")
	}

	printHost(srv)

	return nil
}

// printHost prints every parameter of the host block. Repeated keys
// (LocalForward, IdentityFile, ...) are grouped under the key's first
// occurrence, keeping file order; inline comments are shown after the value.
func printHost(h *host.Host) {
	if h.Comment != "" {
		fmt.Printf("%s\r\n", ansi.Color("# "+h.Comment, "cyan"))
	}

	var order []string
	groups := make(map[string][]host.Param)
	width := 0
	for _, p := range h.Params {
		k := strings.ToLower(p.Key)
		if _, ok := groups[k]; !ok {
			order = append(order, k)
			if len(p.Key) > width {
				width = len(p.Key)
			}
		}
		groups[k] = append(groups[k], p)
	}

	for _, k := range order {
		for i, p := range groups[k] {
			label := ""
			if i == 0 {
				label = p.Key + ":"
			}
			line := ansi.Color(fmt.Sprintf("%-*s", width+1, label), "green") + " " + ansi.Color(p.Value, "white")
			if p.Comment != "" {
				line += "  " + ansi.Color("# "+p.Comment, "cyan")
			}
			fmt.Printf("%s\r\n", line)
		}
	}
}
