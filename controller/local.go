package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/madebyais/bucket/config"
	"github.com/madebyais/bucket/service"
)

// Local controller
type Local struct{}

// DownloadFile is used to download a file
func (l *Local) DownloadFile(c echo.Context) error {
	localFolderPath := c.Get(config.LOCAL_FOLDER_PATH_KEY).(string)
	bucket := c.Param("bucket")
	filename := c.Param("filename")

	filepath := fmt.Sprintf("%s/%s/%s", localFolderPath, bucket, filename)
	return c.File(filepath)
}

// UploadFile is used to upload a file to specific bucket
func (l *Local) UploadFile(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "failed to upload file, error=" + err.Error(),
		})
	}

	localFolderPath := c.Get(config.LOCAL_FOLDER_PATH_KEY).(string)
	bucket := c.Param("bucket")

	localService := service.NewLocal()
	errSaveFile := localService.SaveFile(localFolderPath, bucket, file)
	if errSaveFile != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"message": "failed to save file during upload, error=" + errSaveFile.Error(),
		})
	}

	downloadURL := fmt.Sprintf("/local/%s/%s", bucket, localService.FormatFilename(file.Filename))

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"data": map[string]string{
			"url": downloadURL,
		},
	})
}
