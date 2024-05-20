package helpers

import (
	"fmt"
	"os"
	"strings"
)

func TransformPathFile(filepath string) string {
	baseURL := os.Getenv("BASE_URL")
	urlPath := strings.ReplaceAll(filepath, `\`, "/")
	urlPath = strings.ReplaceAll(urlPath, " ", "-")
	urlPath = strings.TrimPrefix(urlPath, "src/public")
	return fmt.Sprintf("%s/public%s", baseURL, urlPath)
}
