package golvl2hw301

import (
	"log"
)

func Config() {
	appHost := "localhost"
	appPort := "8080"
	log.Printf("App start on %s:%s address!\n", appHost, appPort)
}
