#!/bin/bash

source scripts/constants/colors.sh

chmod +x scripts/docker.sh
chmod +x scripts/prepare.sh
chmod +x scripts/run.sh

echo "Starting docker.sh..."
./scripts/docker.sh
if [ $? -ne 0 ]; then
  echo -e "${RED}Docker setup failed${NC}"
  exit 1
fi

echo "Starting prepare.sh..."
./scripts/prepare.sh

if [ $? -eq 0 ]; then
  echo "${GREEN}prepare.sh completed with code 0${NC}"
else
  echo "${RED}prepare.sh completed with code 1${NC}"
  exit 1
fi

echo "Starting run.sh..."
./scripts/run.sh

if [ $? -eq 0 ]; then
  echo "${GREEN}run.sh completed with code 0${NC}"
else
  echo "${RED}run.sh completed with code 1${NC}"
  exit 1
fi