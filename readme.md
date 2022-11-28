# Measuring DNS resolution latency in a continuous loopback with logging capabilities

The script is written in Golang, it can be run with the following command, or it can be compiled on the platform when needs to run.

```
# Running standalone script
go run dnstest.go

# Building platform specific binary
go build dnstest.go
```

The script is accepting command line arguments for the hostnamea and delay between tests. It can be run as follows:

```
dnstest 2 <hostname1> <hostname2> <hostname3> ...
# Example: dnstes 2 google.com microsoft.com udemy.com
```

It will display the A records of the domain name, IP address associated with that domain, specific resolve time for that hostname and also global resolve time for all the provided domains. For any failure in DNS resolution, it will write the failed domain name along with the timestamp in a local file dnsreq.log.
