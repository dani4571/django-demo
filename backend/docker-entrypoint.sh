#!/bin/bash
echo "HERE?"
cd /app
manage.py migrate
exec gunicorn --bind '[::]:80' --worker-tmp-dir /dev/shm --workers "${GUNICORN_WORKERS:-3}" backend.wsgi:application