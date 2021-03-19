#!/usr/bin/env python

from setuptools import setup, find_packages

setup(name='django-todo-react',
      version='0.0',
      packages=find_packages(),
      scripts=['manage.py'],
      install_requires=['django', 'django-cors-headers', 'djangorestframework']
     )
