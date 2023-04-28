FROM ubuntu:22.04

RUN apt update -y
RUN apt install openssh-server -y
RUN useradd -m hub


CMD service ssh restart && chmod 600 /home/hub/.ssh/authorized_keys && sleep infinity
