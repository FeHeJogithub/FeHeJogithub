import requests

def download_image_to_mnt(url: str, save_path: str) -> None:
    response = requests.get(url, stream=True)
    response.raise_for_status()
    with open(save_path, "wb") as f:
        for chunk in response.iter_content(chunk_size=8192):
            if chunk:
                f.write(chunk)

if __name__ == "__main__":
    IMAGE_URL = "https://plus.unsplash.com/premium_photo-1666672388644-2d99f3feb9f1?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D"
    SAVE_PATH = "imagen.jpg"
    
    try:
        download_image_to_mnt(IMAGE_URL, SAVE_PATH)
        print(f"✅ Imagen descargada correctamente en {SAVE_PATH}")
    except Exception as e:
        print(f"❌ No se pudo descargar la imagen: {e}")