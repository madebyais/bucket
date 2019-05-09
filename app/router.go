package app

import (
	"github.com/madebyais/bucket/controller"

	"github.com/labstack/echo"
)

// NewRouter setup all router for bucket server
func NewRouter(e *echo.Echo) {
	localCtrl := &controller.Local{}
	e.GET("/local/:bucket/:filename", localCtrl.DownloadFile)
	e.POST("/local/:bucket", localCtrl.UploadFile)
}
