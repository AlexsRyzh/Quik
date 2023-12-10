package image

import (
	"github.com/labstack/echo/v4"
	"github.com/quik/backend/internal/api/file"
	"net/http"
)

func Upload(c echo.Context) error {

	filename := c.Param("filename")
	if !file.IsExistFile(filename) {
		return c.String(http.StatusNotFound, "Файл не найден")
	}

	return c.File(file.FilePath + filename)

}
