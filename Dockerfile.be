# https://blog.ploetzli.ch/2020/efficient-multi-stage-build-django-docker/
# Dockerfile
FROM python:slim as common-base

# TODO: Find better versions
RUN apt-get update -y && \
    apt-get install -y libmariadb-dev-compat libmariadb-dev

# ENV DJANGO_SETTINGS_MODULE foo.settings
FROM common-base as base-builder
RUN pip install -U pip setuptools
RUN mkdir -p /app
WORKDIR /app

# Stage 1: Extract dependency information from setup.py alone
# Allows docker caching until setup.py changes
FROM base-builder as dependencies
COPY backend/setup.py .
RUN python setup.py egg_info



# Stage 2: Install dependencies based on the information extracted from the previous step
# Caveat: Expects an empty line between base dependencies and extras, doesn't install extras
# Also installs gunicon in the same step
FROM base-builder as builder
RUN apt-get update && apt-get install -y build-essential python3-dev
RUN mkdir -p /install
COPY --from=dependencies /app/django_todo_react.egg-info/requires.txt /tmp/
RUN sh -c 'pip install --no-warn-script-location --prefix=/install $(grep -e ^$ -m 1 -B 9999 /tmp/requires.txt) gunicorn'
# Everything up to here should be fully cacheable unless dependencies change
# Now copy the application code
COPY backend .


# Stage 3: Install application
RUN sh -c 'pip install --no-warn-script-location --prefix=/install .'
# Stage 4: Install application into a temporary container, in order to have both source and compiled files
# Compile static assets
FROM builder as static-builder
RUN cp -r /install/* /usr/local
RUN sh -c 'python manage.py collectstatic --no-input'


# Stage 5: Install compiled static assets and support files into clean image
FROM common-base
RUN mkdir -p /app
# TODO: Try without this
ENV DYLD_LIBRARY_PATH=/usr/local/mysql/lib/
COPY backend/docker-entrypoint.sh /app/
COPY --from=builder /install /usr/local
COPY --from=static-builder /app/static.dist /app/static.dist
ENTRYPOINT ["/app/docker-entrypoint.sh"]