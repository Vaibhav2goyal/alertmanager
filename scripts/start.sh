#!/bin/sh

set -a


    echo "Starting the Alertmanager - $(date)"

    source /app/.env
    
    /app/alertmanager
