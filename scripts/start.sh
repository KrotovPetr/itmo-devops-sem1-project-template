#!/bin/bash

source scripts/constants/colors.sh

chmod +x scripts/prepare.sh
chmod +x scripts/run.sh

echo "${YELLOW}Starting prepare.sh...${NC}"
./scripts/prepare.sh

if [ $? -eq 0 ]; then
  echo "${GREEN}prepare.sh completed with code 0${NC}"
else
  echo "${RED}prepare.sh completed with code 1${NC}"
  exit 1
fi

echo "${YELLOW}Starting run.sh...${NC}"
./scripts/run.sh

if [ $? -eq 0 ]; then
  echo "${GREEN}run.sh completed with code 0${NC}"
else
  echo "${RED}run.sh completed with code 1${NC}"
  exit 1
fi