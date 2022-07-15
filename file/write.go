package File

import (
	"bufio"
	"log"
	"os"
)

func FileWrite(path []string) {
	file, err := os.OpenFile("file_url.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := bufio.NewWriter(file)
	num := 0
	for i, _ := range path {
		buf.WriteString(path[i])
		buf.WriteByte('\n')
		if num > 10000 {
			buf.Flush()
			num = 0
		}
		num++
	}
	buf.Flush()
}
