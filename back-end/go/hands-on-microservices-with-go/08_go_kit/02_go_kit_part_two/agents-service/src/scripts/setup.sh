#/!bin/bash

#SET HOME FOLDER
#CHANGE THIS TO YOUR HOME FOLDER
YOUR_HOME="/home/paulo"

#YOUR TARGET OS
GOOS="linux"

#SET PATH TO THIS FOLDER
PATH_TO_VIDEO_FOLDER=$YOUR_HOME"/Workspace/web-development/back-end/go/hands-on-microservices-with-go/08_go_kit/02_go_kit_part_two/"

#STOP DOCKER COMPOSE IF IT's RUNNING
sudo docker-compose down

#CLEAN UP
./clean-up.sh

#GENERATE RUNNABLES
#AGENTS
cd $PATH_TO_VIDEO_FOLDER"agents-service/src/api/"
rm ../../main
go build -a -o ../../main main.go

cd $PATH_TO_VIDEO_FOLDER"agents-service/src/scripts"
#RUN DOCKER COMPOSE
sudo docker-compose up -d --build