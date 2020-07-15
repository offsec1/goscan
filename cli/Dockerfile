FROM kalilinux/kali-rolling AS build

# Update
RUN apt-get update && apt-get install -y golang make git zip libc6-dev-i386 nmap dnsrecon dnsenum wfuzz nikto dirb sqlmap hydra

# Setup GO Env
ENV GOROOT=/usr/lib/go
ENV GOPATH=/go
ENV PATH=$PATH:/go/bin

# Setup workdir
WORKDIR /go/src/github.com/marco-lancini/goscan

# Setup project
#COPY goscan/ /go/src/github.com/marco-lancini/goscan
#RUN CGO_ENABLED=0 go build -o /bin/goscan

#FROM scratch
#COPY --from=build /bin/goscan /bin/goscan
#ENTRYPOINT ["/bin/goscan"]
