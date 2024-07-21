package main

import "fpi/create" 
import (
	"os"
	"log"
	"flag"
)

func main(){
	commands := os.Args[1:]

	if len(commands) < 1{
		log.Fatal("Insufficient arguments")
	}


	var overwrite bool

	flag.BoolVar(&overwrite, "overwrite", false, `
		Defines whether to overwrite contents 
		of existing files. True if yes,
		False if no
	`)
	flag.Parse()

	fileNames := flag.Args()

	// log.Println(fileNames)
	create.Create([]string{"routers","controllers","schemas","models","seeders", "services"}, fileNames, overwrite)
}