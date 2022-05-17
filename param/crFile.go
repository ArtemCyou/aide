package param

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CreateFile(argPathCreate string) {
	nameExtension := []string{".txt", ".doc", ".zip"}

	print("Введите максимальное количество файлов которые нужно создать: ")
	askInt, err := bufio.NewReader(os.Stdin).ReadString('\n')
	//askInt, err := reader.ReadString('\n')
	pikaFatal(err)

	askIntNoSpace := strings.TrimSpace(askInt)

	x, err := strconv.Atoi(askIntNoSpace)
	pikaFatal(err)

	for _, name := range nameExtension {
		for i := 1; i <= x; i++ {
			f, err := os.Create(argPathCreate+"/"+strconv.Itoa(i) + name)
			pikaFatal(err)
			fmt.Println("file with name "+f.Name(), "was created")
			f.Close()
		}
	}
}

func pikaFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
