import requests
import hashlib
SERVER = "http://localhost:1234/"
def store(chunk, digest, local=True, server=SERVER):
    if not local:
        dist = "distributed/"
    else:
        dist = ""
    return requests.post(server + dist + "store", data={'key': digest, 'data': chunk}).json()

def findvalue(digest, local=True, server=SERVER):
    if not local:
        dist = "distributed/"
    else:
        dist = ""
    resp = requests.get(server + dist + "find-value/" + digest).json()
    print resp
    if resp['status'] == 'OK':
        return resp['data']
    else:
        return None

def digest_for_chunk(chunk):
    return hashlib.sha1(chunk).hexdigest()
