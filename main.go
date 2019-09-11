package main

import (
	"fmt"
	"github.com/joho/godotenv"
    "log"
	"os"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
func main(){
	initEnv()
	initFrameword()
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
}

func initFrameword(){
	e := echo.New()
	e.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize: 1 << 10, //1KB
	}))
	e.Use(middleware.CORS())
	e.Use(middleware.RequestID())
	e.Use(middleware.GzipWithConfig(middleware.GzipConfig{
		Level: 5,
	}))

	port := os.Getenv("PORT")
	fmt.Printf("Server listening at 3000")
	err := e.Start(":" + port)
	if err != nil {
		fmt.Println(err)
	}
}

func addController() {
}