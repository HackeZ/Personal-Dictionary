FROM golang
MAINTAINER HackerZ
RUN mkdir /PD
ADD main /PD
ADD conf /PD/conf
ADD views /PD/views
ADD static /PD/static
WORKDIR  /PD
EXPOSE 8081
ENTRYPOINT  /PD/main

