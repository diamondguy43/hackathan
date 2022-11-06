#!/usr/bin/env python
#
# Convert webp to jpg
#


import os
from PIL import Image

if __name__ == "__main__":
    baseDir = "../cmpgrounds-pics-webp"
    baseSaveJPG = "../cmpgrounds-pics-jpg"
    baseSavePNG = "../cmpgrounds-pics-png"
    
    # Convert webp to jpg
    for f in os.listdir(f"{baseDir}"):
        webp = f"{baseDir}/{f}"
        rootFilename = f.split('.')[0]

        im = Image.open(webp).convert("RGB")
        im.save(f"{baseSaveJPG}/{rootFilename}.jpg", 'jpeg')

    # Convert jpg to png
    for f in os.listdir(f"{baseDir}"):
        webp = f"{baseDir}/{f}"
        rootFilename = f.split('.')[0]

        im = Image.open(webp).convert("RGB")
        im.save(f"{baseSavePNG}/{rootFilename}.png", 'png')


    
