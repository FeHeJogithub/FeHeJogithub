import requests

if __name__ == "__main__":


    client_id = "2e7e45237055bde0f770"

    client_secret = "f346523b0d62d34daf97da9b29f4cab5e77b1dfb"

    code = "3ded3c8e8659e40347b1"

    access_token = "546442c72cab41236fe7b2f573c0abb698648e87"

                   

    
    url= "https://api.github.com/user/repo"
    payload = {"name" : "git_api_cf_comunidad" }
    headers = {"Accept" : "application/jsonn", "Authorization " : "token " + access_token}
     
    response = requests.post(url, headers=headers, json=payload)
    if response.status_code == 200:
        print( response.json() )
        
                
    else:
        print(response.content)