#!/bin/bash

# Copyright (C) 2016 Yoshiki Shibata. All rights reserved.

go build -o a.out
go build gopl.io/ch1/fetch
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 > w3.xml
./a.out < w3.xml > w3.xmltree
# .\fetch.exe http://www.w3.org/TR/2006/REC-xml11-20060816 | .\e7.18.exe > .\w3.xmltree
# .\fetch.exe http://www.w3.org/TR/2006/REC-xml11-20060816 > w3.xml
# type .\w3.xml | .\e7.18.exe > .\w3.xmltree
