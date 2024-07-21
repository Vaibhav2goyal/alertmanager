#!/bin/sh

set -a


echo "Starting the Alertmanager - $(date)"
##Source the variables
source /app/.env
##Start the application
/app/alertmanager
