package main

import (
    "os"
    "log"
    "elrng/app"
)

func main() {
    app.RunServer(app.AppConf{
        Host:       "localhost",
        Port:       "6969",
        DbUrl:      os.Getenv("DB_URL"),
        JwtSecret:  os.Getenv("JWT_SECRET"),
        Logger:     log.Default(),
    })
}

