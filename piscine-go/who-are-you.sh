#! /bin/bash
./who-are-you.sh | cat -e
curl https://zero.academie.one/assets/superhero/all.json | jq '.[] | select( .id == 70) | .name'