package param

import (
	"fmt"
	"log"
	"os"
)

func CreateListFile(argPathList string)  {
	//create file access.log
	f, err := os.OpenFile(argPathList+"/"+"access.log", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	// read direction
	files, err := os.ReadDir(argPathList)
	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		if _, err := f.WriteString(file.Name() + "\n"); err != nil {
			f.Close() // ignore error; Write error takes precedence
			log.Fatal(err)
		}
		fmt.Println(file.Name())
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
