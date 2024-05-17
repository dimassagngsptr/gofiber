package helpers

import "gofiber/src/configs"

func CountData(table string) int64 {
	var result int64
	configs.DB.Table(table).Where("deleted_at IS NULL").Count(&result)
	return result
}
