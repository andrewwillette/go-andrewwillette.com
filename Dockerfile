FROM ubuntu:latest

ENV DEBIAN_FRONTEND noninteractive

RUN apt update && apt install -y git golang nodejs npm

RUN git clone https://github.com/andrewwillette/go-andrewwillette.com.git 

RUN cd go-andrewwillette/scripts

