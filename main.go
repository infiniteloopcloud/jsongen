package main

import (
	"log"

	"github.com/infiniteloopcloud/jsongen/pkg/updater"
)

// https://go.googlesource.com/go/+archive/refs/tags/go1.20.1/src/encoding/json.tar.gz

var version = "go1.20.1"
var destination = "../json_try"

func main() {
	source, err := updater.DownloadTarGz(version)
	if err != nil {
		log.Fatal(err) // TODO
	}
	err = updater.ExtractTarGz(source, destination)
	if err != nil {
		log.Fatal(err) // TODO
	}

	node, fset, err := updater.ParseAndModify(destination + "/encode.go")
	if err != nil {
		log.Fatal(err) // TODO
	}

	if err := updater.PersistChanges(node, fset, destination+"/encode.go"); err != nil {
		log.Fatal(err) // TODO
		 
	}

}
