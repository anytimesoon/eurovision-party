package frontend

import (
	"embed"
)

//go:generate npm i
//go:generate npm run build
//go:embed all:build
var files embed.FS
