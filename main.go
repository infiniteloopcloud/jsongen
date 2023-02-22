package main

import (
	"flag"
	"github.com/infiniteloopcloud/jsongen/pkg/logger"
	"github.com/infiniteloopcloud/jsongen/pkg/updater"
)

var version = "go1.20.1"
var destination = "../json_try"

func init() {
	flag.StringVar(&version, "version", "go1.20.1", "Version you want to apply the generation (format: go1.20.1)")
	flag.StringVar(&destination, "destination", "../json", "Location of the repository where you want to generate")
	flag.Parse()
}

func main() {
	logger.Info("Start to generate json generation for " + version + " to " + destination)
	logger.Info("Cleanup - remove the go files from the destination")
	err := updater.Clean(destination)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Download Go's tar.gz for the specified version")
	source, err := updater.DownloadTarGz(version)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Extract the tar.gz")
	err = updater.ExtractTarGz(source, destination)
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Analyze AST and apply changes on the encode.go")
	node, fset, err := updater.ParseAndModify(destination + "/encode.go")
	if err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Persist the previous changes on the encode.go")
	if err := updater.PersistChanges(node, fset, destination+"/encode.go"); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Create interfaces.go")
	if err := updater.CreateInterface(destination + "/interfaces.go"); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Create interfaces_test.go")
	if err := updater.CreateInterfaceTest(destination + "/interfaces_test.go"); err != nil {
		logger.Error(err.Error())
	}

	logger.Info("Generation was successful")
}
