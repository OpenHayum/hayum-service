package util

import (
	"io"
	"log"
	"os"

	"bitbucket.org/hayum/hayum-service/config"
)

func InitLogger() {
	logpath := config.App.GetString("logpath")
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.Printf("Using log path: %s\n", logpath)
}
