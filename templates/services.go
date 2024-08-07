package templates

var serviceTemplate string

func ServiceTemplate() string{

	serviceTemplate = `
class %sService:
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

	return serviceTemplate
}