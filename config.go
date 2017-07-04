package main

import (
  "github.com/kelseyhightower/envconfig"
  "github.com/inconshreveable/log15"
  "os"
)

var config = struct {
  Host               string `default:"localhost"`
  Port               int    `default:"8000"`
  DefaultContentType string `default:"application/json"`
}{}

func init(){
  err := envconfig.Process("api", &config)
  if err != nil {
    log15.Crit(err.Error())
    os.Exit(1)
  }
}
