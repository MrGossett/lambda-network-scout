lambda-network-scout
===

Quick little POC to sniff out how requests from Lambda to DynamoDB are routed
when the Lambda is not in a customer-owned VPC.

```
START RequestId: ... Version: $LATEST
2021/08/11 19:25:21 DNS Info: {Addrs:[{IP:52.94.4.90 Zone:}] Err:<nil> Coalesced:false}
2021/08/11 19:25:21 Got Conn: {Conn:0xc00013ae00 Reused:false WasIdle:false IdleTime:0s}
END RequestId: ...
REPORT RequestId: ...	Duration: 23.49 ms	Billed Duration: 24 ms	Memory Size: 512 MB	Max Memory Used: 30 MB
```

`52.94.4.90` is part of the block `52.94.4.0/24`, which is allocated for AWS services in us-east-2.

Now with VPC config with a route to a VPC-E, and public IPs.

```
START RequestId: ... Version: $LATEST
2021/08/11 19:44:48 DNS Info: {Addrs:[{IP:52.94.4.102 Zone:}] Err:<nil> Coalesced:false}
2021/08/11 19:44:48 Got Conn: {Conn:0xc0000de380 Reused:false WasIdle:false IdleTime:0s}
END RequestId: ...
REPORT RequestId: ...	Duration: 29.23 ms	Billed Duration: 30 ms	Memory Size: 512 MB	Max Memory Used: 41 MB
```

Looks exactly the same. Also, disabled the "automatically assign Public IP addresses" config in the VPC subnets, and there was no change.

## Build

```sh
cd cmd/scout
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o scout .
zip function.zip scout
```

## Deploy

Upload `function.zip` as the deployment package. Set the handler to `scout`.
Ensure the function execution role includes `dynamodb:ListTables`.
