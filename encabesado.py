import requests
import json

if __name__ == "__main__":

    url = "https://httpbin.org/post"
    #args = {"nombre": "Fernando", "curso": "python", "nivel": "intermedio" }
    payload = {"nombre": "Fernando", "curso": "python", "nivel": "intermedio" }
    headers = {"Content-Type" : "application/json", "access-token":"12345" }
   # response = requests.post(url, params=args)
  # response = requests.post(url, data=payload)
    response = requests.post(url, data=json.dumps(payload), headers=headers)
    
    #json post se encarga de serializarlos
    # data entonces nosotros nos encargamos de serializarlos
    print(response.url)
    
    
    if response.status_code == 200:
        #print(response.content)
        headers_response = response.headers
        server = headers_response["server"]
        print(server)