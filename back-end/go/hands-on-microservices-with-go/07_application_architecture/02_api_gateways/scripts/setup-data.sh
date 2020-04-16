#!/bin/bash

#SET HOME FOLDER
#CHANGE THIS TO YOUR HOME FOLDER
YOUR_HOME="/home/paulo"

#SET PATH TO THIS FOLDER
PATH_TO_VIDEO_FOLDER=$YOUR_HOME"/Workspace/web-development/back-end/go/hands-on-microservices-with-go/07_application_architecture/02_api_gateways/"

sudo docker volume rm mariadb-agents
sudo docker volume create mariadb-agents
sudo docker run -it --rm -v mariadb-agents:/volume -v $PATH_TO_VIDEO_FOLDER/data:/backup alpine \
    sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/mariadb-agents.tar.bz2"

sudo docker volume rm mariadb-videos
sudo docker volume create mariadb-videos
sudo docker run -it --rm -v mariadb-videos:/volume -v $PATH_TO_VIDEO_FOLDER/data:/backup alpine \
    sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/mariadb-videos.tar.bz2"

sudo docker volume rm mariadb-users
sudo docker volume create mariadb-users
sudo docker run -it --rm -v mariadb-users:/volume -v $PATH_TO_VIDEO_FOLDER/data:/backup alpine \
    sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/mariadb-users.tar.bz2"

sudo docker volume rm postgres-wta
sudo docker volume create postgres-wta
sudo docker run -it --rm -v postgres-wta:/volume -v $PATH_TO_VIDEO_FOLDER/data:/backup alpine \
    sh -c "rm -rf /volume/* /volume/..?* /volume/.[!.]* ; tar -C /volume/ -xjf /backup/postgres-wta.tar.bz2"
