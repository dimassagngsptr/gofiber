package helpers

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadFile(file *multipart.FileHeader) string {
	destination := "src/public"
	epochTime := time.Now().Unix()
	convertFileName := strings.ReplaceAll(file.Filename, " ", "-")
	filePath := filepath.Join(destination, fmt.Sprintf("%d_%s", epochTime, convertFileName))
	if _, err := os.Stat(destination); os.IsNotExist(err) {
		os.Mkdir(destination, os.ModePerm)
	}
	return filePath
}

func ValidateFile() map[string]interface{} {
	functions := make(map[string]interface{})

	functions["SizeUploadValidation"] = func(fileSize int64, maxFileSize int64) error {
		if fileSize > maxFileSize {
			return fiber.NewError(fiber.StatusRequestEntityTooLarge, "Ukuran file melebihi 2MB")
		}
		return nil
	}
	functions["TypeUploadValidation"] = func(buffer []byte, validFileTypes []string) error {
		fileType := http.DetectContentType(buffer)
		isValidFile := functions["isValidFileType"].(func([]string, string) bool)
		if !isValidFile(validFileTypes, fileType) {
			return fiber.NewError(fiber.StatusBadRequest, "Tipe file tidak valid. Hanya png, jpg, jpeg, dan pdf yang diperbolehkan.")
		}
		return nil
	}
	functions["isValidFileType"] = func(validFileTypes []string, fileType string) bool {
		for _, validType := range validFileTypes {
			if validType == fileType {
				return true
			}
		}
		return false
	}
	return functions
}
