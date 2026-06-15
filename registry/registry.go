package registry

import "embed"

//go:embed all:commands all:rules all:templates all:workflows
var FS embed.FS
