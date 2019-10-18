#!/usr/bin/env python3
# -*- coding: utf-8 -*-
#! author = ixsec
from translate import Translator
import io
import sys
def main(argv):
    file_name = argv[0]
    files = io.open(file_name,encoding="utf-8")
    files2 = io.open(file_name+".translate",'a+',encoding="utf-8")
    text_lines = files.readlines()
    for i in text_lines:
        translator = Translator(to_lang="zh")
        translation = translator.translate(i)
        translation.encode('utf-8')
        if len(translation) > 0:
            files2.write(translation+"\n")
    files.close()
    files2.close()
	
if __name__ == "__main__":
    main(sys.argv[1:])
    print("success")