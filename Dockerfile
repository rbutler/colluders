FROM ubuntu:trusty
MAINTAINER Ryan Butler <ryan@rbutler.io>

RUN apt-get update

ADD target/colluders /bin/

