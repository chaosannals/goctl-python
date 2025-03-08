import httpx
from .base import *
from .message import *

class DemoApi:
    def __init__(self, host, port=80, scheme='http'):
        self.__base_url = f'{scheme}://{host}:{port}'

    
    def get_demo(self, req: Request) -> Response:
        r = httpx.get(req.get_path(f'{self.__base_url}/from/:name'))
        if r.status_code != httpx.codes.OK:
            raise DemoApiException(r.content, r.status_code)
        
        result = Response()
        result.set_body(r.json())
        return result
        
    