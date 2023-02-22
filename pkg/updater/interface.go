package updater

import (
	"fmt"
	"os"
	"strings"

	"github.com/infiniteloopcloud/jsongen/pkg/logger"
)

func CreateInterface(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	_, err = fmt.Fprint(file, interfaceGo)
	return err
}

func CreateInterfaceTest(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			logger.Error(err.Error())
		}
	}()

	_, err = fmt.Fprint(file, strings.Join(interfaceTestGo, "\n"))
	return err
}
