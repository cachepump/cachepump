# CachePump
[![Go Report Card](https://goreportcard.com/badge/github.com/cachepump/cachepump)](https://goreportcard.com/report/github.com/cachepump/cachepump)
[![Test status](https://github.com/cachepump/cachepump/actions/workflows/pr_check.yml/badge.svg)](https://github.com/cachepump/cachepump/actions/workflows/pr_check.yml)
[![Go.Dev reference](https://img.shields.io/badge/go.dev-reference-blue?logo=go&logoColor=white)](https://pkg.go.dev/github.com/cachepump/cachepump)
[![Test coverage](https://raw.githubusercontent.com/cachepump/cachepump/master/.github/labels/test_coverage.svg)](https://github.com/cachepump/cachepump/actions/workflows/pr_check.yml)

## About

### What is a CachePump 🤔 ?
CachePump is a smart in-memory cache with a data delivery mechanism.  
For example let's look on next case:  
You have a backend service with http endpoint. Business logic of this service is a very difficult and all requests to service process to long. Results for each requests have some expirity, it may by 10 seconds, 1 houer or more.
If data in service response is no different on big time interval you can use CachePump as view layer.

### How it is work 👨‍💻 ?
CachePump has two important parts:
1. **Internal in-memory storage** which gives fast access to data by unique key.
2. **Scheduler** which deliveries data for each unique key from source defined into configuration file.

You need to define list of sources in configuration file. For eache dataset of source you write request, update interval and unique key. This key will use in request to CachePump.
Scheduler in according with you configuration will receive datasets from sources and saves them to in-memory storage. You can get data by key with using http interface of CachePump.

## Installation
```bash
# Receive source code
git clone https://github.com/cachepump/cachepump.git ./cachepump
cd ./cachepump

# Select version
VERSION='v1.0.0'
git checkout "${VERSION}" -b "branch-${VERSION}"

# Building
go build -o cachepump

# Start
./cachepump -e '0.0.0.0:8080' -l 'INFO' -c './config.yml'
```

## Configuration
You can define sources and way of data delivery in file yaml, see file `./config.yml` in this repository for example. 

### Static data
Static data it is a constant value. For example you can use this source if you need to create mock for http service in little time.
```bash
# Create configuration file.
cat <<EOF > ./my_config.yml
version: '1.0'
sources:
  hello:
    rule: '* * * * * *'
    static: 
      value: '{"data":"Hello world!"}'
  haelth_check:
    rule: '* * * * * *'
    static: 
      value: '{"status":200}'
EOF

# Start CachePump.
./cachepump -e '0.0.0.0:8080' -l 'INFO' -c './my_config.yml'

# Receiving data.
curl -X GET 'http://0.0.0.0:8080/?key=hello'
# {"data":"Hello world!"}
curl -X GET 'http://0.0.0.0:8080/?key=haelth_check'
# {"status":200}
```

### Data from http endpoint
CachePump can receive data from any http resource and cache them. For example you can create a data cached proxy for ClickHouse in little time.
```bash
# Create configuration file.
cat <<EOF > ./my_config.yml
version: '1.0'
sources:
  count_202103:
    rule: '0 */2 * * * *'
    http:
      endpoint: http://0.0.0.0:8123
      method: POST
      header:
      auth:
        user: admin
        password: 'adminadmin'
      body: >
        SELECT date, count(*) 
        FROM DB.Raw_Data
        PREWHERE toYYYYMM(date) = 202103
        GROUP BY date
        ORDER BY date
EOF

# Start CachePump.
./cachepump -e '0.0.0.0:8080' -l 'INFO' -c './my_config.yml'

# Receiving data.
curl -X GET 'http://0.0.0.0:8080/?key=count_202103'
# 673936523
```

### Data from file
You can delivery data to CachePump from static file. For example if you ETL service update a report in static file by scheduler and need to give access to this file by http protocol.
```bash
# Create configuration file.
cat <<EOF > ./my_config.yml
version: '1.0'
sources:
  report_from_elt:
    rule: '0 */2 * * * *'
    file:
      path: my_report_file.csv
EOF

# Start CachePump.
./cachepump -e '0.0.0.0:8080' -l 'INFO' -c './my_config.yml'

# Receiving data.
curl -X GET 'http://0.0.0.0:8080/?key=report_from_elt'
# date,user,count_events
# 2022-01-01,user1,10
# 2022-01-01,user2,14
# 2022-01-03,user6,120
# ...
```