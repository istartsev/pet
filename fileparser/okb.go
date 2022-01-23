package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


func readFileOSOpen(path string) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open file: %v", err)
	}
	defer f.Close()

	buf := make([]byte, 0)
	sentence := ""
	for {
		n, err := f.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		strChunk := strings.ReplaceAll(string(buf[:n]), "\n", "")
		dotIdx := strings.Index(strChunk, ".")
		if dotIdx == -1 {
			sentence += strChunk
		} else {
			sentence += strChunk[:dotIdx+1]
			log.Print(strings.Trim(sentence, " "))
			sentence = strChunk[dotIdx+1:]
		}
	}
}

func readFileIoutils(path string) {
	ioutil.ReadFile(path)
}

func main() {
	s := "123"
	b := []byte(s)
	b[1]='4'

	fmt.Println(s)
	fmt.Println(b)


	readFileOSOpen("test.txt")

}
