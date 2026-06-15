package registry

import "embed"

//go:embed all:skills all:rules all:templates all:workflows .sheave.toml
var FS embed.FS
