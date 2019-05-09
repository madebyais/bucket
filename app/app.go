package app

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/middleware"

	"github.com/labstack/echo"
	"github.com/madebyais/bucket/config"
)

var bucketlogo = `
___                        ___                 
|  |                       |  |            ____   
|  |___  ____  ___ _______ |  |___ _______ |  |___
|   _  | |  |_|  | |   __| |  __ / |  \__| |   ___|
|______/ |_______| |_____| |__|\_\ |_____| |____/
`

// IApp is an interface for App package
type IApp interface {
	Start() error
}

// App represents app package
type App struct {
	server     *echo.Echo
	bucketAuth map[string]string

	Port           string
	ConfigFilePath string
}

// New initiates new application
func New(port string, configFilePath string) IApp {
	return &App{
		bucketAuth:     make(map[string]string),
		Port:           port,
		ConfigFilePath: configFilePath,
	}
}

func (a *App) initServer() {
	a.server = echo.New()
	a.server.HideBanner = true
	a.server.Debug = true

	a.server.Use(middleware.Logger())
	a.server.Use(middleware.BodyLimit("10MB"))
	a.server.Use(middleware.CORS())
	a.server.Use(middleware.Gzip())
	a.server.Use(middleware.Recover())
	a.server.Use(middleware.Secure())
}

func (a *App) initBucketAuthorization(localFolderPath string, bucketName string, bucketToken string) {
	a.bucketAuth[bucketName] = bucketToken

	a.server.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			bucket := c.Param("bucket")

			var token string
			if c.Request().Method == "GET" {
				token = c.QueryParam("token")
			} else {
				token = c.Request().Header.Get("x-bucket-token")
			}

			val, ok := a.bucketAuth[bucket]
			if !ok || val != token {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"status":  "error",
					"message": "unauthorized access",
				})
			}

			c.Set(config.LOCAL_FOLDER_PATH_KEY, localFolderPath)

			return next(c)
		}
	})
}

func (a *App) initStatic(bucketFolderPath string) {
	a.server.Static("/test", bucketFolderPath)
}

func (a *App) initBucketConfig() error {
	bucketConf, err := config.New(a.ConfigFilePath)
	if err != nil {
		return err
	}

	localBucketFolder := bucketConf.Local.Folder
	err = config.InitLocalBucketFolder(localBucketFolder)
	if err != nil {
		return err
	}

	for _, item := range bucketConf.Local.Bucket {
		bucketFolderPath := fmt.Sprintf("%s/%s", localBucketFolder, item.Name)
		config.InitLocalBucketFolder(bucketFolderPath)
		a.initStatic(bucketFolderPath)
		a.initBucketAuthorization(localBucketFolder, item.Name, item.Token)
	}

	return nil
}

// Start is used to start bucket server
func (a *App) Start() error {
	a.initServer()

	err := a.initBucketConfig()
	if err != nil {
		return err
	}

	NewRouter(a.server)
	fmt.Println(bucketlogo)
	return a.server.Start(":" + a.Port)
}
