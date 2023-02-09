import os
import sys

#===This (sys.path.append) is needed for VisualStudioCode
sys.path.append(os.getcwd())

class cOOP:

    def __init__(self):
        #self.set_msgs(msgs)
        print("...in constructor of class cOOP()")


    def pwd_of_this_class(self):
        PWD = os.path.realpath(os.path.dirname(__file__))
        print("e.g. this file's location : " + PWD)
        return PWD

# ============================================================
# == Main
# ============================================================

if __name__ == '__main__':
    res = cOOP()
    res.pwd_of_this_class()

