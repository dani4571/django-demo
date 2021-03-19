FROM centos:8 as builder

RUN dnf install -y nodejs
RUN dnf install -y python3

COPY . /app

WORKDIR /app


# Install dependencies for frontend 
RUN cd frontend && \
    npm install

# Build the prod artifacts
RUN npm run build

# Install dependencies for backend 
RUN pip3 install pipenv && \
    ln -s /usr/bin/python3 /usr/bin/python && \
    pipenv install --system --deploy && \
    pip3 install django-cors-headers && \
    pip3 install djangorestframework

FROM nginx:slim as frontend

COPY --from=builder build /var/www/html/frontend
COPY nginx.conf /etc/nginx/nginx.conf


# https://semaphoreci.com/community/tutorials/dockerizing-a-python-django-web-application
FROM python3:slim as backend



# Set user to non-root user
USER 1001