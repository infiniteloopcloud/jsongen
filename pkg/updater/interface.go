package updater

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateInterface(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("[ERROR] ", err.Error())
		}
	}()

	_, err = fmt.Fprint(file, interfaceGo)
	return err
}

func CreateInterfaceTest(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Println("[ERROR] ", err.Error())
		}
	}()

	_, err = fmt.Fprint(file, strings.Join(interfaceTestGo, "\n"))
	return err
}
