to build run

`docker build -t goback:latest -f Dockerfile.goback .` 

from directory above this one

to run in kubernetes run

`kubectl apply -f manifests/goback_kube.yml`

image is published here https://hub.docker.com/repository/docker/danielcurran90/gostonks
and can be pulled using 

`docker pull danielcurran90/gostonks`

you will just have to change the image name (or just change references here)