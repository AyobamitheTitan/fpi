package main

import (
	"fmt"
	"os"
	"log"
)

func main(){
	pwd, err := os.Getwd()

	if err != nil{
		log.Fatalf("%s", err)
	}

	fmt.Println(pwd)
}