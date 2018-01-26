#!/usr/bin/python3
import binascii
import base64

def main():
    receivedString = '49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d'
    expectedString = 'SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t'

    message = binascii.unhexlify(receivedString)

    decoded = base64.b64encode(message).decode('ascii')

    if (decoded == expectedString):
        print(message)
    else:
        print("ERROR: No Match!")

if __name__ == '__main__':main()
