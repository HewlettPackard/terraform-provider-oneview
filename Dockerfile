FROM golang:alpine
MAINTAINER "Priyanka Sood <priyanka.sood@hpe.com>" 

ENV TERRAFORM_VERSION=0.13.6

ENV USER root
RUN mkdir -p /usr/local/terraform
RUN apk update && \
    apk add curl jq python3 bash ca-certificates git openssl unzip wget && \
    cd /tmp && \
    wget https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip && \
unzip terraform_${TERRAFORM_VERSION}_linux_amd64.zip -d /usr/local/terraform/

ENV GOROOT /usr/local/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin
WORKDIR /go/src/github.com/HewlettPackard/terraform-provider-oneview

COPY . /go/src/github.com/HewlettPackard/terraform-provider-oneview

RUN cd $GOPATH/src/github.com/HewlettPackard/terraform-provider-oneview

RUN go get github.com/kardianos/govendor
ENV PATH $PATH:/usr/local/terraform/
