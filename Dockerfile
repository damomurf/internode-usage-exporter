FROM golang:1.6-alpine
MAINTAINER damian@murf.org

COPY ./bin/internode-usage-exporter /

EXPOSE 9099

CMD /internode-usage-exporter
