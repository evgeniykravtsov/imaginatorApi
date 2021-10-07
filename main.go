package main

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	e := echo.New()

	//Проверяем наличие папки images, при ее отсутствии создаем её
	existsDir("./images")

	e.POST("/saveImage", save)
	e.GET("/getImageById", getImage)
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
	existsDir(fmt.Sprintf("./images/" + id.String()))
	path:= fmt.Sprintf("images/" + id.String() + "/"  + image.Filename)
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

func getImage (c echo.Context) error {
	id := c.QueryParam("id")
	dirs, err := ioutil.ReadDir("./images/" + id)
	if err != nil {
		return c.String(http.StatusBadRequest, "We have not image for this id!")
	}

	if len(dirs) == 1 {
		fileName := dirs[0].Name()
		return c.File("./images/" + id + "/" + fileName)
	} else {
		return c.String(http.StatusBadRequest, "OK")
	}
}

func existsDir(path string)  {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
	}
}