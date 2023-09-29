#! /bin/bash

find . -type f -name "*.sh" | rev | cut -d'/' -f1 | rev | cut -d'.' -f1 | sort -r
