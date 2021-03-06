// This server can run on Google App Engine.
package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/cybozu-go/log"
	"github.com/cybozu-go/neco/gcp"
	"github.com/cybozu-go/neco/gcp/app"
	"google.golang.org/appengine"
	"gopkg.in/yaml.v2"
)

const (
	cfgFile = ".necogcp.yml"
)

func loadConfig() (*gcp.Config, error) {
	cfg, err := gcp.NewConfig()
	if err != nil {
		return nil, err
	}
	f, err := os.Open(cfgFile)
	if err != nil {
		// If cfgFile does not exist, use neco-test config
		return gcp.NecoTestConfig(), nil
	}
	defer f.Close()
	err = yaml.NewDecoder(f).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func main() {
	// seed math/random
	rand.Seed(time.Now().UnixNano())

	cfg, err := loadConfig()
	if err != nil {
		log.ErrorExit(err)
	}

	server, err := app.NewServer(cfg)
	if err != nil {
		log.ErrorExit(err)
	}
	http.HandleFunc("/shutdown", server.HandleShutdown)

	appengine.Main()
}
