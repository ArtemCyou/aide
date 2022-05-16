package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

func main() {

	argPathClean := flag.String("clean", ".", "Очистить целевой каталог")
	argPathGroup := flag.String("dir", ".", "Dir для группировки файлов")
	flag.Parse()

	switch os.Args[1] {
	case "-clean":
		print(*argPathClean)
		cleanDir(*argPathClean)
	case "-dir":
		fmt.Println(*argPathGroup)
		groupFiles(*argPathGroup)
	default:
		os.Exit(1)
	}

}

func cleanDir(argPathClean string) {
	pathDir, err := os.ReadDir(argPathClean)
	pikaFatal(err)
	for _, f := range pathDir {
		os.RemoveAll(path.Join([]string{argPathClean, f.Name()}...))
	}
}

func groupFiles(argPathGroup string) {
	filesExtension := []string{
		".txt",
		".doc",
		".docx",
		".pdf",
		".xlsx",
		".bmp",
		".jpg",
		".rtf",
		".pptx",
		".conf",
		".cfg",
		".net",
		".deny",
		".allow",
		".exe",
	}
	files, err := os.ReadDir(argPathGroup)
	pikaFatal(err)

	for _, file := range files {
		for _, ext := range filesExtension {
			if strings.HasSuffix(file.Name(), ext) {
				newDirPath := argPathGroup + "ALL" + strings.ToUpper(ext)
				if _, err := os.Stat(newDirPath); os.IsNotExist(err) {
					err := os.Mkdir(newDirPath, 0644)
					pikaFatal(err)
				}
				err := os.Rename(argPathGroup+file.Name(), newDirPath+"/"+file.Name())
				pikaFatal(err)
			}
		}
	}
}

func pikaFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
