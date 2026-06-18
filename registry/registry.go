package registry

import "embed"

//go:embed all:skills all:rules all:templates all:workflows all:settings .sheave.toml
var FS embed.FS
