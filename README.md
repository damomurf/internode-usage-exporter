# internode-usage-exporter


## Configuration

Create a YAML configuration file, for example at ```/etc/internode-usage/config.yaml```:

```
---
username: <internode username>
password: <internode password>
```

## Run with Docker

docker run -d -p 9099:9099 -v /etc/internode-usage:/config quay.io/damomurf/internode-usage-exporter /internode-usage-exporter -config /config/config.yaml

[![Build Status](https://ci.murf.org/api/badges/damian/internode-usage-exporter/status.svg)](https://ci.murf.org/damian/internode-usage-exporter)
[![Docker Repository on Quay](https://quay.io/repository/damomurf/internode-usage-exporter/status?token=23aa85f4-3700-4a9e-8925-882b59e5c652 "Docker Repository on Quay")](https://quay.io/repository/damomurf/internode-usage-exporter)

