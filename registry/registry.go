package registry

import "embed"

//go:embed all:commands all:rules all:templates all:workflows .sheave.toml
var FS embed.FS
