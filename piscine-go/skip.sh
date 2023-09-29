#! /bin/bash

ls -l | sort -k9 | awk 'NR % 2 == 0'