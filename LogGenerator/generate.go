package main

import (
	"flag"
	"fmt"
	"github.com/bxcodec/faker"
	"math/rand"
	"os"
	"strconv"
)

func usage() {
	fmt.Printf("Usage: %s [OPTIONS] [files] [lines]\n", os.Args[0])
	fmt.Println("files is the number of files to generate.")
	fmt.Println("lines is the number of lines per file.")
	flag.PrintDefaults()
	os.Exit(1)
}

func genLine() string {
	methods := [4]string{"GET", "POST", "PUT", "DELETE"}
	index := rand.Intn(len(methods))
	return fmt.Sprintf("%s %s %s [%s] %s", faker.IPv4(), faker.Username(), faker.Email(), faker.Timestamp(), methods[index])
}

func generateLogFiles(numberOfFiles, linesPerFile int, prefix string) {
	fmt.Println("Generating", numberOfFiles, "files with", linesPerFile, "lines in each.")

	for i := 0; i < numberOfFiles; i++ {
		fileName := prefix + strconv.Itoa(i) + ".log"
		f, _ := os.Create(fileName)
		for i := 0; i < linesPerFile; i++ {
			f.WriteString(genLine() + "\n")
		}
		f.Close()
	}
}

func main() {
	var prefix string
	flag.Usage = usage
	flag.StringVar(&prefix, "prefix", "gen_", "The prefix to use for generated files.")
	flag.Parse()

	if flag.NArg() != 2 {
		flag.Usage()
	}
	numberOfFiles, err := strconv.Atoi(os.Args[1])
	if err != nil {
		flag.Usage()
	}

	linesPerFile, err := strconv.Atoi(os.Args[2])
	if err != nil {
		flag.Usage()
	}

	generateLogFiles(numberOfFiles, linesPerFile, prefix)

}
