FROM centurylink/ca-certs
MAINTAINER damian@murf.org

COPY ./bin/internode-usage-exporter /

EXPOSE 9099

CMD /internode-usage-exporter
