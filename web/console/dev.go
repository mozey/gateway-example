// +build dev

package console

import (
	"github.com/mozey/gateway/internal/config"
	"net/http"
	"path"
)

// VFS contains project assets.
var VFS http.FileSystem

func init() {
	conf := config.New()
	VFS = http.Dir(path.Join(conf.Dir(), "web", "console", "data"))
}


