package template

func init() {
	WebProjectFiles["api/ping.go"] = `// Generate By Template
package api

import (
	"github.com/chaimch/akita-inu/formutil"
	"github.com/gin-gonic/gin"

	"{{ .ProjectPath }}/handler"
)

func Ping(c *gin.Context) {
	formutil.Handle(&handler.PingForm{}, c)
}
`
}
