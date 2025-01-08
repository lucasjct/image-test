import requests

def sendRequest():
    request = requests.get('https://www.google.com')
    print(request.status_code)

sendRequest()