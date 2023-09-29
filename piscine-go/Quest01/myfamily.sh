#! /bin/bash
curl https://zero.academie.one/assets/superhero/all.json | jq --argjson heroid "$HERO_ID" '.[] | select( .id == $heroid ) | .connections | .relatives' | sed 's/"//g'
