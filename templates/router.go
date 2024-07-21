package templates

var routerTemplate string

func RouterTemplate() string{
	routerTemplate = `
from fastapi import APIRouter

%s_router = APIRouter(prefix="")
controller = %sController()


%s_router.get("")\
	(controller.get)


%s_router.get("/{ID}")\
	(controller.get_by_id)




%s_router.post("")\
	(controller.store)


	

%s_router.patch("/{ID}")\
	(controller.update)



%s_router.delete("/{ID}")\
	(controller.delete)
`

	return routerTemplate
}