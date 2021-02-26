FROM ubuntu:latest

RUN apt-get update
COPY bin/ha-hello-world .
RUN chmod +x ha-hello-world

USER nobody

CMD ./ha-hello-world