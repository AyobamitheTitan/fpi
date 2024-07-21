package templates

var modelTemplate string 


func ModelTemplate () string{
	modelTemplate = `
class %s():
	pass
`
	return modelTemplate
}