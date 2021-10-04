package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	e := echo.New()

	//Проверяем наличие папки images, при ее отсутствии создаем её
	existsDir("./images")

	e.POST("/saveImage", save)
	e.Logger.Fatal(e.Start(":1323"))
}

func save (c echo.Context) error {
	//Ищем картинку по полю image
	image, err := c.FormFile("image")

	if err != nil {
		return c.String(http.StatusBadRequest, "FormData dont have image!")
	}

	src, errImage := image.Open()
	if errImage != nil {
		return c.String(http.StatusBadRequest, "Can not open image!")
	}

	defer src.Close()

	id := uuid.New()
	fileTypeSlice := strings.Split(image.Filename, ".")
	var fileType string
	if len(fileTypeSlice) == 2 {
		fileType = fmt.Sprintf("." + fileTypeSlice[1] )
	}
	path:= fmt.Sprintf("images/" + id.String()  + fileType)
	dst, errCreate := os.Create(path)

	if errCreate != nil {
		return c.String(http.StatusBadRequest, "Can not save image!")
	}
	defer dst.Close()

	// Copy
	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.String(http.StatusOK, "Image save! " + "ID is "+ id.String())
}

func existsDir(path string)  {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}