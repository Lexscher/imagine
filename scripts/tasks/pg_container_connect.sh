#!/bin/bash

docker exec -it $(docker ps --filter name=pg --format "{{.Names}}") /bin/bash