#! /bin/bash

curl https://zero.academie.one/assets/superhero/all.json | jq -r '.[] | select( .id == 170) | .name'
curl https://zero.academie.one/assets/superhero/all.json | jq '.[] | select( .id == 170) | .powerstats | .power'
curl https://zero.academie.one/assets/superhero/all.json | jq -r '.[] | select( .id == 170) | .appearance | .gender'
