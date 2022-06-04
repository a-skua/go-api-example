FROM debian

WORKDIR /home/e2e

RUN apt-get update && \
    apt-get -y install curl jq

RUN useradd e2e

USER e2e
