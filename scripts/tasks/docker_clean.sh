#!/bin/bash

# remove all containers
docker container rm -f $(docker container ps -a -q)

# clear volumes
docker volume prune -f