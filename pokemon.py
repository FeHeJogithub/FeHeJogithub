import requests


def get_pokemons(url= "http://pokeapi.co/api/v2/pokemon-form/", offset=0):
    
#def get_datos(url= "EgZjaHJvbWUqCggBEAAYsQMYgAQyDAgAEEUYORixAxiABDIKCAEQABixAxiABDIHCAIQABiABDIHCAMQABiABDIHCAQQABiABDIHCAUQABiABDINCAYQLhivARjHARiABDIKCAcQABixAxiABDIHCAgQABiABDIHCAkQABiABNIBCDU2MDRqMGo3qAIIsAIB8QWrJaWlkGgq0PEFqyWlpZBoKtA", offset=0):
    args = {"offset":offset} if offset else {}
    
    response = requests.get(url, params=args)
    
    if response.status_code==200:
       
       payload=response.json()
       results = payload.get("results", [])
       
       if results:
           for pokemon in results:
               name=pokemon["name"]
               print(name)
               #print(url)
    #raw_input=True la cambie por input       
    next = input("Continuar listando? [Y/N]").lower()
    if next =="y":
       get_pokemons(offset=offset+200)
                   
if __name__ == "__main__":
    
   # url= "EgZjaHJvbWUqCggBEAAYsQMYgAQyDAgAEEUYORixAxiABDIKCAEQABixAxiABDIHCAIQABiABDIHCAMQABiABDIHCAQQABiABDIHCAUQABiABDINCAYQLhivARjHARiABDIKCAcQABixAxiABDIHCAgQABiABDIHCAkQABiABNIBCDU2MDRqMGo3qAIIsAIB8QWrJaWlkGgq0PEFqyWlpZBoKtA"
    url= "http://pokeapi.co/api/v2/pokemon-form/"
    get_pokemons()