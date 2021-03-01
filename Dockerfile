FROM ubuntu:latest

RUN apt-get -y update && apt-get -y upgrade
COPY bin/ha-hello-world .
RUN chmod +x ha-hello-world

USER nobody

CMD ./ha-hello-world