#!/bin/bash

source scripts/constants/colors.sh
source scripts/constants/path.sh


if [ ! -f $APP_BIN_PATH ]; then
  echo -e "${RED}Application not compiled or not found at $APP_BIN_PATH${NC}"
  exit 1
fi

echo -e "Starting application..."
./$COMPILE_TO &
if [ $? -ne 0 ]; then
  echo -e "${RED}Error starting application${NC}"
  exit 1
fi

echo -e "${GREEN}Application started${NC}"

exit 0