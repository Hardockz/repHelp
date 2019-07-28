#!/bin/bash

cd ~/repHelp && make && docker build -t rephelp -f deploy/Dockerfile . && cd ~/repHelp/deploy && docker-compose up


