package app 

import (
    "net/http"
    "elrng/routes"
)

func addRoutes(mux *http.ServeMux) {
    routes.AddUserRoutes(mux)
}

