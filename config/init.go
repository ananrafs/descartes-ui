package config

import (
	"io"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

var (
	Configuration Config
	once          sync.Once

	serverLoc = "./config/files/server.yaml"
)

func init() {
	once.Do(func() {
		file, err := os.Open(serverLoc)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		byt, err := io.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		err = yaml.Unmarshal(byt, &Configuration)
		if err != nil {
			log.Fatal(err)
		}
	})
}
