package routes 

import (
    "net/http"
    "database/sql"
    "fmt"
    "elrng/data"
)

// Example
func getUsers() http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        if user, err := data.UserDao.GetById(req.Context(), 1); err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "", http.StatusNotFound)
            } else {
                msg := fmt.Sprintf("UserDao.GetById: %s\n", err.Error())
                http.Error(w, msg, http.StatusInternalServerError)
            }
        } else {
            fmt.Println(user)
            w.Write([]byte("OK"))
        }
    })
}

func AddUserRoutes(mux *http.ServeMux) {
    mux.Handle("/api/v1/users", getUsers())
}

