import requests


if __name__ == "__main__":
  # url = "https://plus.unsplash.com/premium_photo-1666672388644-2d99f3feb9f1?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
    url = " https://plus.unsplash.com/premium_photo-1666669764052-4f8175f051dc?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
    response = requests.get(url, stream=True) #Realiza la peticion sin descargar el contenido
    with open("image.jpg", "wb") as file:
        for chunk in response.iter_content():#Descarga el contenido poco a poco
            file.write(chunk)
  
    response.close()