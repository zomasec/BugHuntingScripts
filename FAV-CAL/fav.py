#!/bin/python3

# The Script is takes from you favicon picture link and convert it as hash to search on shodan on it like:  http.favicon.hash:1031234597

import mmh3
import requests
import codecs
import argparse

def get_arguments():
    parser = argparse.ArgumentParser(description="Process The Inputs")
    parser.add_argument("-l", "--link", dest="link", help="Link of The Favicon Picture To Convert it to Hash.")
    options = parser.parse_args()
    if not options.link:
        parser.error("[-] Please specify Link Path of The Favicon Picture.")

    return options



Favicon_Link = get_arguments().link
response = requests.get(Favicon_Link)
favicon = codecs.encode(response.content,"base64")
hash = mmh3.hash(favicon)
print(hash)

# Credits : @bug4you
# Usage: ./fav.py -l https://nasa.gov/favicon.ico
