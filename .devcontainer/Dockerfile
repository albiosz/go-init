FROM ubuntu:22.04
ARG UBUNTU_VERSION_CODENAME=jammy
ARG GO_VERSION=1.21.1
ARG LIQUIBASE_VERSION=4.23.2

RUN apt update && \
  apt -y upgrade && \
  apt -y install git && \
  apt -y install make && \
  apt -y install wget && \
  apt -y install gnupg

# install postgresql client 15
RUN sh -c 'echo "deb http://apt.postgresql.org/pub/repos/apt ${UBUNTU_VERSION_CODENAME}-pgdg main" > /etc/apt/sources.list.d/pgdg.list' && \
 wget --quiet -O - https://www.postgresql.org/media/keys/ACCC4CF8.asc | apt-key add -

RUN apt update && \
  apt -y install postgresql-client-15

# install go
RUN wget https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz

ENV GOPATH="/home/dev/go"
ENV GOBIN="/home/dev/go/bin"
ENV PATH="/usr/local/go/bin:${GOBIN}:${PATH}"

RUN adduser dev && \
  usermod -aG sudo dev

USER dev

WORKDIR /home/dev/app

RUN go install github.com/cosmtrek/air@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install golang.org/x/tools/gopls@latest && \
  go install golang.org/x/tools/cmd/goimports@latest && \
  go install honnef.co/go/tools/cmd/staticcheck@latest
