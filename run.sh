#!/bin/bash

cd ~/mattermost-bot-test && make && docker build -t rephelp -f deploy/Dockerfile . && cd ~/mattermost-bot-test/deploy && docker-compose up


