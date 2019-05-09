package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

// ILocal represents interface for local service
type ILocal interface {
	FormatFilename(filename string) string
	SaveFile(localFolderPath string, bucket string, file *multipart.FileHeader) error
}

// Local represents struct for local service
type Local struct{}

// NewLocal initiates new local service
func NewLocal() ILocal {
	return &Local{}
}

// FormatFilename returns new filename by replacing spaces with dash
func (l *Local) FormatFilename(filename string) string {
	return strings.ReplaceAll(filename, " ", "-")
}

// SaveFile is used to save file to specific bucket
func (l *Local) SaveFile(localFolderPath string, bucket string, file *multipart.FileHeader) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	// Destination
	filepath := fmt.Sprintf("%s/%s/%s", localFolderPath, bucket, l.FormatFilename(file.Filename))
	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return nil
}
