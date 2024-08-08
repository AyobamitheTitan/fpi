package templates

var schemaTemplate string

func SchemaTemplate () string{
	schemaTemplate = `
from typing import Optional
from pydantic import BaseModel

class %sCreate(BaseModel):
	pass



class %sUpdate(BaseModel):
	pass
`
	return schemaTemplate
}