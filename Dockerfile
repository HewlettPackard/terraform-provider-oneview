FROM golang:alpine

RUN apk add --update git make sudo unzip wget openssh

RUN wget --no-check-certificate https://releases.hashicorp.com/terraform/0.6.16/terraform_0.6.16_linux_amd64.zip && \
    unzip terraform_0.6.16_linux_amd64.zip -d /usr/local/terraform && \
    rm terraform_0.6.16_linux_amd64.zip && \
    rm /usr/local/terraform/terraform-provider-* && \
    rm /usr/local/terraform/terraform-provisioner-chef

RUN mkdir -p /go/src/HewlettPackard/ && \
    cd /go/src/HewlettPackard && \
    git clone https://github.com/HewlettPackard/terraform-provider-oneview.git