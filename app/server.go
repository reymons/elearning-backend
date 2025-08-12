package app

import (
    "log"
    "net"
    "net/http"
    "elrng/db"
)

type AppConf struct {
    Host        string
    Port        string
    JwtSecret   string
    DbUrl       string 
    Logger      *log.Logger
}

func RunServer(conf AppConf) {
    if err := db.Initialize(conf.DbUrl); err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    addr := net.JoinHostPort(conf.Host, conf.Port)
    mux := http.NewServeMux()
    addRoutes(mux)
    handler := addMiddlewares(mux, conf.Logger)

    log.Printf("Listening on %s\n", addr)
    log.Fatal(http.ListenAndServe(addr, handler))
}
