package main

import (
	"encoding/json"
	"fmt"
	dh "github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/handler"
	"github.com/andikanugraha11/golang-boilerplate-awesome-echo/app/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	// color variables (in bytecode form)
	green   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109})
	white   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109})
	yellow  = string([]byte{27, 91, 57, 48, 59, 52, 51, 109})
	red     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109})
	blue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109})
	magenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109})
	cyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109})
	reset   = string([]byte{27, 91, 48, 109})

)

func main()  {
	// Apps Settings
	var appMode string
	CLIConf := os.Args

	e := echo.New()
	e.HideBanner = true

	if len(CLIConf) > 1 {
		switch CLIConf[1] {
		case "DEV":
			appMode = "DEV"
			if CLIConf[2] == "debug" {
				e.Debug = true
			}
		case "PROD":
			appMode = "PRODUCTION"
		default:
			fmt.Println("fill first argument with PROD or DEV to set application mode")
			os.Exit(0)
		}
	} else {
		appMode = "DEV"
	}

	name := fmt.Sprintf("EchoAwesome-%s", appMode)

	// ENV Config
	switch appMode {
	case "DEV":
		// Env File
		viper.SetConfigFile("dev.env")

		// LOGGING
		e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			var reqMethod string
			var resStatus int
			var statusColor, methodColor, resetColor string
			// request and response object
			req := c.Request()
			res := c.Response()
			// rendering variables for response status and request method
			resStatus = res.Status
			reqMethod = req.Method// for response status
			switch {
			case resStatus >= http.StatusOK && resStatus < http.StatusMultipleChoices:
				statusColor = green
			case resStatus >= http.StatusMultipleChoices && resStatus < http.StatusBadRequest:
				statusColor = white
			case resStatus >= http.StatusBadRequest && resStatus < http.StatusInternalServerError:
				statusColor = yellow
			default:
				statusColor = red
			}
			// for request method
			switch reqMethod {
			case "GET":
				methodColor = blue
			case "POST":
				methodColor = cyan
			case "PUT":
				methodColor = yellow
			case "DELETE":
				methodColor = red
			case "PATCH":
				methodColor = green
			case "HEAD":
				methodColor = magenta
			case "OPTIONS":
				methodColor = white
			default:
				methodColor = reset
			}
			// reset to return to the normal terminal color variables (kinda default)
			resetColor = reset
			// print formatting the custom logger tailored for DEVELOPMENT environment
			fmt.Printf("\n[%s] %v |%s %3d %s| %8s | %10s |%s %-7s %s %s",
				name, // name of server (APP) with the environment
				time.Now().Format("2006/01/02 - 15:04:05"), // TIMESTAMP for route access
				statusColor, resStatus, resetColor, // response status
				req.Proto,                          // protocol
				c.RealIP(),                         // client IP
				methodColor, reqMethod, resetColor, // request method
				req.URL, // request URI (path)
			)
		}))
	default:
		// Env File
		viper.SetConfigFile("prod.env")

		// LOGGING
		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Format: fmt.Sprintf("\n[%s] | ${host} | ${time_custom} | ${status} | ${latency_human} | ${remote_ip} | ${bytes_in} bytes_in | ${bytes_out} bytes_out | ${method} | ${uri} ",
				name,
			),
			CustomTimeFormat: "2006/01/02 15:04:05", // custom readable time format
			Output:           os.Stdout,             // output method
		}))
	}

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	// Database Settings
	db, err := config.DBConnection()
	if err != nil {
		log.Panicf("Terjadi masalah pada koneksi database. %s\n", err.Error())
	}

	dHandler := dh.NewDevHandler(db)
	// API LIST
	APIRoutes(e, db)
	DevRotes(e, dHandler)

	// stores routes available in the system in a JSON file
	data, err := json.MarshalIndent(e.Routes(), "", "  ")
	if err != nil {
		fmt.Println(err) // error handling in Golang works this way
	}
	ioutil.WriteFile("./routes.json", data, 0644)

	e.Logger.Fatal(e.Start(":1323"))
}
