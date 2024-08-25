import os
import base64

encoded = "aWQ="  # base64 coded 'id'
decoded = base64.b64decode(encoded).decode("utf-8")
os.system(decoded)
