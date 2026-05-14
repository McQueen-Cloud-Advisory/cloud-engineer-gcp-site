package main

import (
    "errors"
    "log"
    "net/http"
    "os"
    "path"
    "strings"
)

const siteRoot = "/site"

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fileServer := http.FileServer(http.Dir(siteRoot))

    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        cleanPath := path.Clean("/" + strings.TrimPrefix(r.URL.Path, "/"))
        fullPath := path.Join(siteRoot, strings.TrimPrefix(cleanPath, "/"))

        info, err := os.Stat(fullPath)
        if err == nil {
            if info.IsDir() {
                indexPath := path.Join(fullPath, "index.html")
                if _, err := os.Stat(indexPath); err != nil {
                    serveNotFound(w, r)
                    return
                }
            }

            fileServer.ServeHTTP(w, r)
            return
        }

        if errors.Is(err, os.ErrNotExist) {
            serveNotFound(w, r)
            return
        }

        http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
    })

    server := &http.Server{
        Addr:    ":" + port,
        Handler: handler,
    }

    log.Printf("serving static site on %s", server.Addr)
    if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
        log.Fatal(err)
    }
}

func serveNotFound(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    http.ServeFile(w, r, path.Join(siteRoot, "404.html"))
}