package src

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
)

var cache = make(map[string]cachedContent)
var cacheExpiration = 1 * time.Hour
var PUBLIC_FOLDER string

type cachedContent struct {
	HTML       string
	Expiration time.Time
}

func listFiles(c echo.Context) error {
	basePath := c.Param("folder")
	filePath := path.Join(PUBLIC_FOLDER, basePath)

	if file, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.String(http.StatusNotFound, "Folder not found")
		} else {
			return c.String(http.StatusInternalServerError, "Failed to read directory")
		}
	} else if !file.IsDir() {
		return c.File(filePath)
	}

	if cached, exists := cache[filePath]; exists && time.Now().Before(cached.Expiration) {
		return c.HTML(http.StatusOK, cached.HTML)
	}

	files, err := readDirectory(filePath)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to read directory")
	}

	var fileLinks []string
	for _, file := range files {
		oneFilePath := path.Join(basePath, file.Name())
		fileLinks = append(fileLinks, fmt.Sprintf("<li><a href=/%s>%s</a></li>", oneFilePath, file.Name()))
	}

	css := `
body {
  font-family: Arial, sans-serif;
  background-color: #f0f0f0;
  margin: 0;
  padding: 0;
}

.container {
  max-width: 800px;
  margin: 0 auto;
  padding: 20px;
  background-color: #fff;
  box-shadow: 0 2px 5px rgba(0, 0, 0, 0.1);
}

h1 {
  color: #333;
}

ul {
  list-style: none;
  padding: 0;
}

li {
  margin-bottom: 8px;
}
`

	html := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
      <style>%s</style>
		</head>
		<body>
			<div class="container">
				<h1>%s</h1>
				<ul>%s</ul>
			</div>
		</body>
		</html>`, css, filePath, strings.Join(fileLinks, ""))

	cache[filePath] = cachedContent{
		HTML:       html,
		Expiration: time.Now().Add(cacheExpiration),
	}

	return c.HTML(http.StatusOK, html)
}

func readDirectory(folder string) ([]os.FileInfo, error) {
	dir, err := os.Open(folder)
	if err != nil {
		return nil, err
	}
	defer dir.Close()

	files, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func Entry() {
	e := echo.New()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	PUBLIC_FOLDER = os.Getenv("PUBLIC_FOLDER")
	if PUBLIC_FOLDER == "" {
		fmt.Println("Need folder environment variable")
		os.Exit(1)
	}

	e.GET("/:folder", listFiles)
	e.GET("/", listFiles)

	e.Logger.Fatal(e.Start(":" + port))
}
