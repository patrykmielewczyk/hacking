require 'base64'

encoded = "aWQ="  # base64 coded 'id'
decoded = Base64.decode64(encoded)

system(decoded)
