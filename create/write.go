package create

import (
	"os"
	"fmt"
	"strings"
	"fpi/templates"
)

func WriteToFile(dirName string, fileName string) error{
	fileCreate := fileName + ".py"

	if dirName == "routers"{
		defaultRouter := fmt.Sprintf(
			templates.RouterTemplate(),
			fileName,
			strings.Title(fileName),
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
			strings.Title(fileName),
		)
		return os.WriteFile(fileCreate, []byte(defaultController), os.FileMode(0755))

	} else if dirName == "models"{
		defaultModel := fmt.Sprintf(
			templates.ModelTemplate(),
			strings.Title(fileName),
		)
		return os.WriteFile(fileCreate, []byte(defaultModel), os.FileMode(0755))

	} else if dirName == "schemas"{
		defaultSchema := fmt.Sprintf(
			templates.SchemaTemplate(),
			strings.Title(fileName),
			strings.Title(fileName),
		)
		return os.WriteFile(fileCreate, []byte(defaultSchema), os.FileMode(0755))

	} else if dirName == "services"{
		defaultService := fmt.Sprintf(
			templates.ServiceTemplate(),
			strings.Title(fileName),
		)
		return os.WriteFile(fileCreate, []byte(defaultService), os.FileMode(0755))
		
	}
	return nil
}