import requests
#import json

if __name__ == "__main__":
   # url = "https://www.google.com.do/"
    
    url = " https://www.instagram.com/ligadelafarandula/"
    response = requests.get(url)
    
    
if response.status_code == 200:
    content = response.content 
    #file = open("google.html", "wb")
    file = open("farandula.html", "wb")
    file.write(content)
    file.close()
  #  print(response.content)
  
    
   