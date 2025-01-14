#!/bin/bash

source scripts/constants/colors.sh
source scripts/constants/dbconfig.sh
source scripts/constants/path.sh


export PGPASSWORD=$DB_PASSWORD

echo -e "Creating prices table..."
psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -c "
CREATE TABLE IF NOT EXISTS prices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    category VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    create_date TIMESTAMP WITH TIME ZONE NOT NULL
);"
if [ $? -ne 0 ]; then
  echo -e "${RED}Error creating prices table${NC}"
  exit 1
fi
echo -e "${GREEN}Prices table created successfully${NC}"

echo -e "Compiling application..."
go build -o $COMPILE_TO $COMPILE_FROM
if [ $? -ne 0 ]; then
  echo -e "${RED}Error compiling application${NC}"
  exit 1
fi
echo -e "${GREEN}Application compiled to $COMPILE_TO${NC}"

exit 0