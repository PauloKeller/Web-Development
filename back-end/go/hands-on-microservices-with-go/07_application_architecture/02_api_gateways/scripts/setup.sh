#/!bin/bash

#SET HOME FOLDER
#CHANGE THIS TO YOUR HOME FOLDER
YOUR_HOME="/home/paulo"

#YOUR TARGET OS
GOOS="linux"

#SET PATH TO THIS FOLDER
PATH_TO_VIDEO_FOLDER=$YOUR_HOME"/Workspace/web-development/back-end/go/hands-on-microservices-with-go/07_application_architecture/02_api_gateways/"

#STOP DOCKER COMPOSE IF IT's RUNNING
sudo docker-compose down

#CLEAN UP
./clean-up.sh

#GENERATE RUNNABLES
#USERS
cd $PATH_TO_VIDEO_FOLDER"users-service/"
go build -a -o main main.go
#MANAGERS
cd $PATH_TO_VIDEO_FOLDER"agents-service/src/"
go build -a -o ../main main.go
#VIDEOS
cd $PATH_TO_VIDEO_FOLDER"videos-service/src/"
go build -a -o ../main main.go
#SESSIONS
cd $PATH_TO_VIDEO_FOLDER"sessions-service/"
go build -a -o main main.go
#WTA
cd $PATH_TO_VIDEO_FOLDER"wta-service/src/"
go build -a -o ../main main.go
#API GATEWAY 1
cd $PATH_TO_VIDEO_FOLDER"api-gateway-1/src/"
go build -a -o ../main main.go
#API GATEWAY 2
cd $PATH_TO_VIDEO_FOLDER"api-gateway-2/src/"
go build -a -o ../main main.go

cd $PATH_TO_VIDEO_FOLDER"scripts"
#RUN DOCKER COMPOSE
sudo docker-compose up --build