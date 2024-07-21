package templates

var controllerTemplate string

func ControllerTemplate() string{
	controllerTemplate = `
class %sController:
	def __init__(self):
		pass

	
	async def find(self, ID):
		pass

	
	async def get(self,):
		pass


	async def get_by_id(self, ID):
		pass


	async def update(self, ID):
		pass


	
	async def store(self, body):
		pass


	
	async def delete(self, ID):
		pass
`
	return controllerTemplate
}