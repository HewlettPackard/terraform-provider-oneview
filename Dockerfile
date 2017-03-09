FROM golang:1.8.0-alpine

RUN apk add --update git make sudo unzip wget openssh

ADD . /go/src/github.com/HewlettPackard/terraform-provider-oneview

ENV GLIDE_VERSION 0.12.1
ENV GLIDE_DOWNLOAD_URL https://github.com/Masterminds/glide/releases/download/v${GLIDE_VERSION}/glide-v${GLIDE_VERSION}-linux-amd64.tar.gz

ENV TERRAFORM_VERSION 0.8.8
ENV TERRAFORM_DOWNLOAD_URL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip

# Install terraform binary
RUN wget --no-check-certificate ${TERRAFORM_DOWNLOAD_URL} -O terraform_linux_amd64.zip && \
    unzip terraform_linux_amd64.zip -d /usr/local/terraform && \
    rm terraform_linux_amd64.zip && \
    find /usr/local/terraform
#    rm /usr/local/terraform/terraform-provider-* && \
#    rm /usr/local/terraform/terraform-provisioner-chef

# Install glide tools
RUN wget --no-check-certificate "$GLIDE_DOWNLOAD_URL" -O glide.tar.gz \
    && tar -xzf glide.tar.gz \
    && mv linux-amd64/glide /usr/bin/ \
    && rm -r linux-amd64 \
    && rm glide.tar.gz

# Build the plugin
RUN mkdir -p /go/src/github.com/HewlettPackard/ && \
    cd /go/src/github.com/HewlettPackard/terraform-provider-oneview && \
    glide install && \
    CGO_ENABLED=0 go build -a -installsuffix cgo -o terraform-provider-oneview && \
    mv terraform-provider-oneview /usr/local/terraform/ && \
    cd /go && \
    rm -rf *
