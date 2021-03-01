FROM ubuntu:latest

RUN apt-get update && apt-get upgrade
COPY bin/ha-hello-world .
RUN chmod +x ha-hello-world

USER nobody

CMD ./ha-hello-world