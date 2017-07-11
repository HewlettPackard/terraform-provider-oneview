FROM alpine

# Add the recently compiled Go binaries
ADD bins/linux/ /usr/local/terraform/

# Add the folder above to the path
ENV PATH $PATH:/usr/local/terraform/
