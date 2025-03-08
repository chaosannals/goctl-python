
class Request:
    def __init__(self):
        self.__path = {}

    @property
    def name(self):
        return self.__path['name']

    @name.setter
    def name(self, v):
        self.__path['name'] = v
    
    
    def get_path(self, path):
        for k, v in self.__path.items():
            path = path.replace(':' + k, v)
        return path

class Response:
    def __init__(self):
        self.__json = {}

    @property
    def message(self):
        return self.__json['message']

    @message.setter
    def message(self, v):
        self.__json['message'] = v
    
    def get_body(self):
        return self.__json
    def set_body(self, v):
        self.__json = v
    