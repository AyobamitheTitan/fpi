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

func AddFiles(dirName string, fileNames []string){
	for _, fileName := range fileNames{
		NewFile(dirName, fileName)
	}
}


func NewDir(dirName string, fileNames []string){
	if _, err := os.Stat(dirName); err == nil{
		AddFiles(dirName, fileNames)
	}else {
		err := os.Mkdir(dirName, os.FileMode(0755))

		if err != nil{
			log.Fatalf("%s", err)
		}

		log.Printf("%s directory created successfully", dirName)

		AddFiles(dirName, fileNames)
	}
}
