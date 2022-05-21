package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"paramDop/param"
)

func main() {

	argPathClean := flag.String("clean", ".", "Очистить целевой каталог")
	argPathGroup := flag.String("dir", ".", "Dir для группировки файлов")
	argPathList := flag.String("list", ".", "Создает документ со списком всех файлов в директории")
	argPathCreate := flag.String("file", ".", "Создает нужное кол-во файлов в указанной директории")
	flag.Parse()

	switch os.Args[1] {
	case "-clean":
		print(*argPathClean)
		cleanDir(*argPathClean)
	case "-dir":
		fmt.Println(*argPathGroup)
		groupFiles(*argPathGroup)
	case "-list":
		fmt.Println(*argPathList)
		param.CreateListFile(*argPathList)
	case "-file":
		fmt.Println(*argPathCreate)
		param.CreateFile(*argPathCreate)
	default:
		os.Exit(1)
	}
}

func cleanDir(argPathClean string) {
	param.AntiDummy(argPathClean)

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
		".zip",
		".rar",
		".mp4",
		".jpeg",
		".mp3",
		".mov",
		".go",
	}
	files, err := os.ReadDir(argPathGroup)
	pikaFatal(err)

	param.AntiDummy(argPathGroup)

	for _, file := range files {
		if strings.Contains(file.Name(), os.Args[0]) {
			continue
		} else {
			for _, ext := range filesExtension {
				if strings.HasSuffix(file.Name(), ext) {
					newDir := argPathGroup + "/" + "ALL" + strings.ToUpper(ext)
					if _, err := os.Stat(newDir); os.IsNotExist(err) {
						err := os.Mkdir(newDir, 0644)
						pikaFatal(err)

					}

					if _, err := os.Stat(newDir + "/" + file.Name()); !os.IsNotExist(err) {
						err := os.Rename(argPathGroup+"/"+file.Name(), newDir+"/"+"new"+file.Name())
						pikaFatal(err)
					} else {
						err := os.Rename(argPathGroup+"/"+file.Name(), newDir+"/"+file.Name())
						pikaFatal(err)
					}
				}
			}
		}
	}
}

func pikaFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
