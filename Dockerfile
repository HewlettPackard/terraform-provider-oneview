FROM golang:1.8.0-alpine

# Define terraform version to download
ARG TERRAFORM_VERSION=0.9.3

# Install unzip, we'll use it later
RUN apk add --no-cache unzip

# Add terraform binary without using wget so we can pass the
# proxy stuff directly using --build-arg HTTP_PROXY=$HTTP_PROXY
ADD https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip /tmp/terraform.zip
RUN unzip /tmp/terraform.zip -d /usr/local/terraform && \
    rm -rf /tmp/terraform.zip && \
    /usr/local/terraform/terraform version

# Add the source code to compile it inside
ADD . /go/src/github.com/HewlettPackard/terraform-provider-oneview

# Build our code inside
RUN cd /go/src/github.com/HewlettPackard/terraform-provider-oneview && \
    CGO_ENABLED=0 go build -a -tags netgo -ldflags '-s -w' -o /usr/local/terraform/terraform-provider-oneview && \
    cd /go && rm -rf *
