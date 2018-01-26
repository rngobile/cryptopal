#!/usr/bin/python3
import binascii

def sxor(s1,s2):
    return ''.join(chr(ord(a) ^ ord(b)) for a,b in zip(s1,s2))

def main():
    string1 = '1c0111001f010100061a024b53535009181c'
    string2 = '686974207468652062756c6c277320657965'
    resultString = '746865206b696420646f6e277420706c6179'

    hex1 = binascii.hexlify(string1)
    hex2 = binascii.hexlify(string2)
    resultHex = binascii.unhexlify(resultString)

    finalString = sxor(hex1,hex2)

    print(finalString)
    print(resultString)

if __name__ == '__main__':main()
