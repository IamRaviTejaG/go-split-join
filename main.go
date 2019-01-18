package main

import (
	"flag"
	"fmt"
	"gosj/sj"
)

func main() {
	option := flag.String("opt", "SPLIT", "The option: SPLIT / JOIN")
	filename := flag.String("file", "file.txt", "The option: SPLIT / JOIN")
	splitsize := flag.Int("size", 0, "Split size, mandatory for SPLITing files (in MB)")
	nParts := flag.Int("parts", 0, "Number of parts, mandatory for JOINing files")
	flag.Parse()
	if *option == "SPLIT" || *option == "split" {
		sj.Split(*filename, *splitsize)
	} else if *option == "JOIN" || *option == "join" {
		sj.Join(*filename, *nParts)
	} else {
		fmt.Println("Error! Invalid value for parameter -opt")
	}
}
