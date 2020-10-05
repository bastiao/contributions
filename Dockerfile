FROM golang:1.14.3

ADD bin/contributions /usr/bin/contributions
ENTRYPOINT ["contributions"]
