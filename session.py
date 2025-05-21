import requests


if __name__ == "__main__":
    url = "https://github.com/FeHeJogithub/FeHeJogithub"
    
    
    session = requests.session()
    session.auth = ( "fernandohernandez@gmail.com", "FeHeJo010101")
    
    response = session.get(url)
    if response.ok:
        print(response.content)
        
    else:
        
        print(response.content)