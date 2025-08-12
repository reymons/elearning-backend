package app 

import (
    "net/http"
    "log"
)

func loggerMiddleware(next http.Handler, logger *log.Logger) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        var (
            method = req.Method
            path   = req.URL.Path
            addr   = req.RemoteAddr
            proto  = req.Proto
        )

        logger.Printf("%s %s %s, %s\n", method, path, proto, addr)
        next.ServeHTTP(w, req)
    })
}

func addMiddlewares(next http.Handler, logger *log.Logger) http.Handler {
    next = loggerMiddleware(next, logger)
    return next
}

