package main

import (
	"fmt"
	"runtime/debug"
	"strings"
)

// Set at build time via -ldflags "-X main.version=... -X main.commit=... -X main.date=...".
// These names match goreleaser's default ldflags.
var (
	version = ""
	commit  = ""
	date    = ""
)

// versionString reports the version injected by ldflags, falling back to the
// build info Go embeds in the binary (module version for `go install pkg@vX`,
// VCS revision/time for builds from a git checkout).
func versionString() string {
	v, c, d := version, commit, date

	if info, ok := debug.ReadBuildInfo(); ok {
		if v == "" && info.Main.Version != "" && info.Main.Version != "(devel)" {
			v = info.Main.Version
		}
		dirty := false
		for _, s := range info.Settings {
			switch s.Key {
			case "vcs.revision":
				if c == "" {
					c = s.Value
					if len(c) > 12 {
						c = c[:12]
					}
				}
			case "vcs.time":
				if d == "" {
					d = s.Value
				}
			case "vcs.modified":
				dirty = s.Value == "true"
			}
		}
		if dirty && commit == "" && c != "" {
			c += "-dirty"
		}
	}

	if v == "" {
		v = "dev"
	}

	var extra []string
	if c != "" {
		extra = append(extra, "commit "+c)
	}
	if d != "" {
		extra = append(extra, "date "+d)
	}
	if len(extra) > 0 {
		return fmt.Sprintf("%s (%s)", v, strings.Join(extra, ", "))
	}
	return v
}
