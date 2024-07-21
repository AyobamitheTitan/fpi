package main

import "fpi/create" 
import (
	"fmt"
	"os"
	"log"
"flag"
)

func main(){
	pwd, err := os.Getwd()

	if err != nil{
		log.Fatalf("%s", err)
	}

	fmt.Println(pwd)
}