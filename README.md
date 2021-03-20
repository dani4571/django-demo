## Introduction

This is a simple Todo application built off Django (including the Django REST Framework for API CRUD operations) and React. For a complete walkthrough, see [Build a To-Do application Using Django and React](https://www.digitalocean.com/community/tutorials/build-a-to-do-application-using-django-and-react/edit)

## Requirements
* Python3
* Pipenv

## Getting started
1. Clone the project to your machine ```[git clone https://github.com/Jordanirabor/django-todo-react]```
2. Navigate into the diretory ```[cd django-todo-react]```
3. Source the virtual environment ```[pipenv shell]```
4. Install the dependencies ```[pipenv install]```
5. Navigate into the frontend directory ```[cd frontend]```
5. Install the dependencies ```[npm install]```

## Run the application
You will need two terminals pointed to the frontend and backend directories to start the servers for this application.

1. Run this command to start the backend server in the ```[backend]``` directory: ```[python manage.py runserver]``` (You have to run this command while you are sourced into the virtual environment)
2. Run this command to start the frontend development server in the ```[frontend]``` directory: ```[npm install]``` (This will start the frontend on the adddress [localhost:3000](http://localhost:3000))

## Built With

* [React](https://reactjs.org) - A progressive JavaScript framework.
* [Python](https://www.python.org/) - A programming language that lets you work quickly and integrate systems more effectively.
* [Django](http://djangoproject.org/) - A high-level Python Web framework that encourages rapid development and clean, pragmatic design.

## Creator Credit

This demo app was originally built for a scotch.io (acquired in 2020 by DigitalOcean) article by [Jordan Irabor](https://github.com/Jordanirabor/django-todo-react)


### Docker info 

#### How to run
* Build frontend ```[docker build -t frontend:latest -f Dockerfile.fe .]```
* Build backend ```[docker build -t backend:latest -f Dockerfile.be .]```
* Run from docker-compose ```[docker-compose up]```

#### TODO
* Clean up the backend dockerfile
* Make config work for sqlite or mysql
* Get command to run manually backend to run manually


Sources:
https://github.com/do-community/django-todo-react.git
https://blog.ploetzli.ch/2020/efficient-multi-stage-build-django-docker/
https://packaging.python.org/guides/distributing-packages-using-setuptools/#configuring-your-project
https://lincolnloop.com/blog/using-setuppy-your-django-project/