package create

import (
	"os"
	"fmt"
	"strings"
	"fpi/templates"
)

func MakePascalCase(word string) string{
	splitWord := strings.Split(word,"_")
	pascalCase := ""

	for _,word := range(splitWord){
		pascalCase += strings.Title(word)
	}
	return pascalCase
}

func WriteToFile(dirName string, fileName string) error{
	fileCreate := fileName + ".py"
	pascalCase := MakePascalCase(fileName)

	if dirName == "routers"{
		defaultRouter := fmt.Sprintf(
			templates.RouterTemplate(),
			fileName,
			pascalCase,
			fileName,
			fileName,
			fileName,
			fileName,
			fileName,
		)
		return os.WriteFile(fileCreate, []byte(defaultRouter), os.FileMode(0755))

	} else if dirName == "controllers"{
		defaultController := fmt.Sprintf(
			templates.ControllerTemplate(),
			pascalCase,
		)
		return os.WriteFile(fileCreate, []byte(defaultController), os.FileMode(0755))

	} else if dirName == "models"{
		defaultModel := fmt.Sprintf(
			templates.ModelTemplate(),
			pascalCase,
		)
		return os.WriteFile(fileCreate, []byte(defaultModel), os.FileMode(0755))

	} else if dirName == "schemas"{
		defaultSchema := fmt.Sprintf(
			templates.SchemaTemplate(),
			pascalCase,
			pascalCase,
		)
		return os.WriteFile(fileCreate, []byte(defaultSchema), os.FileMode(0755))

	} else if dirName == "services"{
		defaultService := fmt.Sprintf(
			templates.ServiceTemplate(),
			pascalCase,
		)
		return os.WriteFile(fileCreate, []byte(defaultService), os.FileMode(0755))
		
	}
	return nil
}