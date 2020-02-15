# kibanaRefreshFields
Refresh kibana index fields via kibanas API
## Expected variables
I have this setup in my ~/.bashrc, another deployment would be to place this in k8s.
```/bin/bash
export KIBANA_USERNAME=jmainguy
export KIBANA_PASSWORD="supersecret"
export KIBANA_URL=logs-lab.example.com
export KIBANA_INDEX="filebeat-*"
# Optional, if you want to filter out some fields
export KIBANA_FILTER="kubernetes.labels.jenkins"
```

## Usage
```/bin/bash
./kibanaRefreshFields
```

## PreBuilt Binaries
Grab Binaries from [The Releases Page](https://github.com/Jmainguy/kibanaRefreshFields/releases)

## Build
```/bin/bash
export GO111MODULE=on
go build
```
