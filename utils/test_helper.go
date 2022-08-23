package utils

import (
	"net/http"
	"testing"

	"github.com/muharamdani/go-base-rest-api/db"

	"github.com/gavv/httpexpect/v2"
	"github.com/gin-gonic/gin"
)

func PrepareTest(t *testing.T, engine *gin.Engine) *httpexpect.Expect {
	db.Connect("test")

	e := httpexpect.WithConfig(httpexpect.Config{
		Client: &http.Client{
			Transport: httpexpect.NewBinder(engine),
			Jar:       httpexpect.NewJar(),
		},
		Reporter: httpexpect.NewAssertReporter(t),
		Printers: []httpexpect.Printer{
			httpexpect.NewDebugPrinter(t, true),
		},
	})

	return e
}

func CleanUp(table string) {
	db.DB.Exec("DELETE FROM " + table)
}