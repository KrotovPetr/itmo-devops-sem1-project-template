#!/bin/bash

source scripts/constants/colors.sh

IMAGE_NAME="itmo-postgres-image"
CONTAINER_NAME="itmo-postgres-container"

if [[ "$(docker images -q $IMAGE_NAME 2> /dev/null)" != "" ]]; then
  echo -e "${YELLOW}Removing existing Docker image...${NC}"
  docker rmi $IMAGE_NAME
  if [ $? -ne 0 ]; then
    echo -e "${RED}Error removing Docker image${NC}"
    exit 1
  fi
fi

echo -e "${YELLOW}Building Docker image...${NC}"
docker build -t $IMAGE_NAME .
if [ $? -ne 0 ]; then
  echo -e "${RED}Error building Docker image${NC}"
  exit 1
fi

if lsof -i:$DB_PORT &>/dev/null; then
  echo -e "${RED}Port $DB_PORT is already in use. Please free the port before proceeding.${NC}"
  exit 1
fi

if [[ "$(docker ps -a -q -f name=$CONTAINER_NAME)" != "" ]]; then
  echo -e "${YELLOW}Removing existing Docker container...${NC}"
  docker rm -f $CONTAINER_NAME
  if [ $? -ne 0 ]; then
    echo -e "${RED}Error removing Docker container${NC}"
    exit 1
  fi
fi

echo -e "${YELLOW}Starting Docker container...${NC}"
docker run --name $CONTAINER_NAME -p $DB_PORT:5432 -d $IMAGE_NAME
if [ $? -ne 0 ]; then
  echo -e "${RED}Error starting Docker container${NC}"
  exit 1
fi

echo -e "${GREEN}Docker container started successfully${NC}"