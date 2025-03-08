from gen import DemoApi, DemoApiException, Request

try:
    demo_api = DemoApi('127.0.0.1', 8888, 'http')
    req = Request()
    req.name = 'me'
    resp = demo_api.get_demo(req)
    print(resp.message)
except DemoApiException as e:
    print(e.message)