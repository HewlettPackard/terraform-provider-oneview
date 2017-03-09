FROM golang:1.8.0-alpine

RUN apk add --update git make sudo unzip wget openssh

ADD . /go/src/github.com/HewlettPackard/terraform-provider-oneview

ENV TERRAFORM_VERSION 0.8.8
ENV TERRAFORM_DOWNLOAD_URL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Add terraform binary without using wget so we can pass the
# proxy stuff directly using --build-arg HTTP_PROXY=$HTTP_PROXY
ADD ${TERRAFORM_DOWNLOAD_URL} /temp/terraform.zip
RUN unzip /temp/terraform.zip -d /usr/local/terraform && \
    rm -rf /temp/terraform.zip && \
    /usr/local/terraform/terraform version

# Build our code inside
RUN cd /go/src/github.com/HewlettPackard/terraform-provider-oneview && \
    CGO_ENABLED=0 go build -a -installsuffix cgo -o /usr/local/terraform/terraform-provider-oneview && \
    cd /go && rm -rf *
