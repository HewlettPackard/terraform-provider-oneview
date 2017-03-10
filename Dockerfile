FROM golang:1.8.0-alpine

RUN apk add --no-cache --update unzip

ADD . /go/src/github.com/HewlettPackard/terraform-provider-oneview

# Define terraform version to download
ENV TERRAFORM_VERSION 0.8.8
ENV TERRAFORM_DOWNLOAD_URL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Add terraform binary without using wget so we can pass the
# proxy stuff directly using --build-arg HTTP_PROXY=$HTTP_PROXY
ADD ${TERRAFORM_DOWNLOAD_URL} /tmp/terraform.zip
RUN unzip /tmp/terraform.zip -d /usr/local/terraform && \
    rm -rf /tmp/terraform.zip && \
    /usr/local/terraform/terraform version

# Build our code inside
RUN cd /go/src/github.com/HewlettPackard/terraform-provider-oneview && \
    CGO_ENABLED=0 go build -a -tags netgo -ldflags '-s -w' -o /usr/local/terraform/terraform-provider-oneview && \
    cd /go && rm -rf *
