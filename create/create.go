package create

import (
	"os"
	"log"
)

var toOverwrite bool

func NewFile(dirName string, fileName string){
	err := os.Chdir(dirName)

	if err != nil{
		log.Fatalf("%s", err)
	}

	fileToCreate := fileName + ".py"
	if _, err := os.Stat(fileToCreate); err == nil && toOverwrite == false{
		os.Chdir("..")
		return
	}

	_, err = os.Create(fileToCreate)

	if err != nil{
		log.Fatalf("%s", err)
	}

	log.Printf("%s created successfully", fileName)
	if err = WriteToFile(dirName, fileName); err != nil{
		log.Fatalf("%s", err)
	}

	err = os.Chdir("..")

	if err != nil{
		log.Fatalf("%s", err)
	}
}
