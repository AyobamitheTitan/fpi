# FPI

Fpi is command line tool to take away the repetitive
task of setting up modules in a FastAPI application.

## Background

Over the course of my internship programme (March to August 2024), I was tasked with using FastAPI to build the backend of various systems I worked on. Over time, setting up a project, creating the routers, models, schemas, controllers and services of each package everytime became quite cumbersome for me. So i decided to build a tool to reduce this load. FPI is the result of this pursuit.

I named it FPI because I built it to help me set up <b>F</b>astA<b>PI</b> applications

## How to use
FPI assumes your code structure is as follows

```
- src
    - mypackage
        - controllers
        - models
        - routers
        - schemas
        - services
main.py
```

To set up your application, cd into mypackage

```
cd mypackage
```

Then run the command

```
fpi tv_channels
```

`tv_channels` in this case is the name of the entity you want to add to your app.

This command results in the following



* /mypackage/controllers/tv_channels.py

```
class TvChannelsController:
    def __init__(self):
        pass
    
    async def find(self, ID):
        pass

    async def store(self, body):
        pass

    async def update(self, ID, body):
        pass

    async def delete(self, ID):
        pass
    
    async def get(self,):
        pass
```


* /mypackage/models/tv_channels.py

(add a link to the documentation to creating models in FastAPI)
```
class TvChannels(Base):
    pass
```




* /mypackage/routers/tv_channels.py

```
from fastapi import APIRouter

tv_channels_router = APIRouter(prefix="/tv_channels")
controller = TvChannelsController()

tv_channels_router.get("")\
    (controller.get)

tv_channels_router.get("/{ID}")\
    (controller.find)

tv_channels_router.post("")\
    (controller.store)

tv_channels_router.patch("/{ID}")\
    (controller.update)

tv_channels_router.delete("/{ID}")\
    (controller.delete)

```




* /mypackage/schemas/tv_channels.py

```
from typing import Optional
from pydantic import BaseModel


class TvChannelsCreate(BaseModel):
    pass

class TvChannelsUpdate(BaseModel):
    pass
```




* /mypackage/services/tv_channels.py

```
class TvChannelsService:
    def __init__(self):
        pass
    
    async def find(self, ID):
        pass

    async def store(self, body):
        pass

    async def update(self, ID, body):
        pass

    async def delete(self, ID):
        pass
    
    async def get(self,):
        pass
```

#### Flags

- overwrite

    If a file of the same name exists in the folder you run `fpi` in,
    you can set the overwrite flag to decide whether or not to overwrite the
    contents of that file.

    By default it's set to false
