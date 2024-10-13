#!/bin/bash

ENV_FILE="../../config/.env"
SCHEMA_FILE="../../models/persistence/sqlite"

DB_USER="myuser"     # replace with your username
DB_NAME="myappdb"           # replace with your database name

# Load environment variables from .env

if [ -f "$ENV_FILE" ]; then
    # Echo the contents of the .env file
    echo "Setting up using the contents of $ENV_FILE:"
    # cat "$ENV_FILE"
    export $(grep -v '^#' "$ENV_FILE" | xargs)
    echo "Args exported! \n"
else
    echo "Error: $ENV_FILE does not exist."
    echo "Create a .env file that has DATABASE values for setup"
fi

# Function to run SQL script and check for errors
run_sql_script() {
    local sql_file=$1

    # Execute the SQL script
    psql -U "$DB_USER" -d "$DB_NAME" -f "$sql_file"

    # Check the exit status
    if [ $? -ne 0 ]; then
        echo "Error: Failed to execute $sql_file."
        echo "Error code: $?"
        exit 1  # Exit the script with an error code
    else
        echo "Successfully executed $sql_file.\n"
    fi
}

# Create the database and user
run_sql_script "create_database.sql"

run_sql_script "../../models/persistence/sqlite/schema.sql"


