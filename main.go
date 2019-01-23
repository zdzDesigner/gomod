package main

import (
	log "github.com/sirupsen/logrus"
	"moddemo/util"
	"moddemo/util/lang"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")

	util.GetName()
	lang.GetLang()
}
