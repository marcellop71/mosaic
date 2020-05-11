import json

import ctypes
from ctypes import c_char_p

libname = 'libmosaic.so'
mosaic = ctypes.cdll.LoadLibrary(libname)

mosaic.newRandomOrg.argtypes = [c_char_p, c_char_p]
mosaic.newRandomOrg.restype = c_char_p

mosaic.newRandomAuth.argtypes = [c_char_p]
mosaic.newRandomAuth.restype = c_char_p

mosaic.newRandomUserkey.argtypes = [c_char_p, c_char_p, c_char_p]
mosaic.newRandomUserkey.restype = c_char_p

mosaic.newRandomSecret.argtypes = [c_char_p, c_char_p]
mosaic.newRandomSecret.restype = c_char_p

mosaic.rewritePolicy.argtypes = [c_char_p]
mosaic.rewritePolicy.restype = c_char_p

mosaic.checkPolicy.argtypes = [c_char_p, c_char_p]
mosaic.checkPolicy.restype = c_char_p

mosaic.authpubsOfPolicy.argtypes = [c_char_p]
mosaic.authpubsOfPolicy.restype = c_char_p

mosaic.policyOfCiphertextJson.argtypes = [c_char_p]
mosaic.policyOfCiphertextJson.restype = c_char_p

mosaic.selectUserAttrs.argtypes = [c_char_p, c_char_p, c_char_p]
mosaic.selectUserAttrs.restype = c_char_p

mosaic.encrypt.argtypes = [c_char_p, c_char_p, c_char_p]
mosaic.encrypt.restype = c_char_p

mosaic.decrypt.argtypes = [c_char_p, c_char_p]
mosaic.decrypt.restype = c_char_p

#mosaic.getPbcCurve.argtypes = [c_char_p, c_char_p]
#mosaic.getPbcCurve.restype = c_char_p

class Mosaic:
    @staticmethod
    def newRandomOrg(lib, curveJson):
        return mosaic.newRandomOrg(lib, curveJson)

    @staticmethod
    def newRandomAuth(orgJson):
        return mosaic.newRandomAuth(orgJson)

    @staticmethod
    def newRandomUserkey(user, attr, authprvJson):
        return mosaic.newRandomUserkey(user, attr, authprvJson)

    @staticmethod
    def newRandomSecret(orgJson, seed):
        return mosaic.newRandomSecret(orgJson, seed)

    @staticmethod
    def rewritePolicy(policy):
        return mosaic.rewritePolicy(policy)

    @staticmethod
    def checkPolicy(policy, userattrsJson):
        return mosaic.checkPolicy(policy, userattrsJson)

    @staticmethod
    def authpubsOfPolicy(policy):
        return mosaic.authpubsOfPolicy(policy)

    @staticmethod
    def policyOfCiphertextJson(ctJson):
        return mosaic.policyOfCiphertextJson(ctJson)

    @staticmethod
    def selectUserAttrs(user, policy, userattrsJson):
        return mosaic.selectUserAttrs(user, policy, userattrsJson)

    @staticmethod
    def encrypt(secret, policy, authpubsJson):
        return mosaic.encrypt(secret, policy, authpubsJson)

    @staticmethod
    def decrypt(ctJson, userattrsJson):
        return mosaic.decrypt(ctJson, userattrsJson)

    #@staticmethod
    #def getPbcCurve(curveStr):
    #    return mosaic.getPbcCurve(curveStr)
