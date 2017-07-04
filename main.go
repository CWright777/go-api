package main

import (
  "fmt"
  "runtime"
  "net/http"
  "time"

  "github.com/CWright777/go-api/endpoint"
  "github.com/inconshreveable/log15"
  "github.com/facebookgo/httpdown"
)

func main() {
  log15.Info(
          "Starting api",
          "go_version", runtime.Version(),
          "go_OS", runtime.GOOS,
          "go_Arch", runtime.GOARCH,
          "host", config.Host,
          "port", config.Port,
  )


  server := &http.Server{
    Addr:    fmt.Sprintf("%s:%d", config.Host, config.Port),
    Handler: endpoint.CreateRouter(),
  }

  hd := &httpdown.HTTP{
    StopTimeout: 10 * time.Second,
    KillTimeout: 1 * time.Second,
  }

  if err := httpdown.ListenAndServe(server, hd); err != nil {
    log15.Crit("serve error", "err", err)
  }

}
