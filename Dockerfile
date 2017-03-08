FROM golang:alpine

RUN apk add --update git make sudo unzip wget openssh

ADD . /go/src/github.com/HewlettPackard/terraform-provider-oneview

RUN wget --no-check-certificate https://releases.hashicorp.com/terraform/0.6.16/terraform_0.6.16_linux_amd64.zip && \
    unzip terraform_0.6.16_linux_amd64.zip -d /usr/local/terraform && \
    rm terraform_0.6.16_linux_amd64.zip && \
    rm /usr/local/terraform/terraform-provider-* && \
    rm /usr/local/terraform/terraform-provisioner-chef

RUN mkdir -p /go/src/github.com/HewlettPackard/ && \
    cd /go/src/github.com/HewlettPackard/terraform-provider-oneview && \
    wget --no-check-certificate https://glide.sh/get -O ./glide.sh && \
    ls -altr; chmod +x ./glide.sh && \
    sh ./glide.sh && \
    glide install && \
    CGO_ENABLED=0 go build -a -installsuffix cgo -o terraform-provider-oneview && \
    mv terraform-provider-oneview /usr/local/terraform/ && \
    cd /go && \
    rm -rf *
