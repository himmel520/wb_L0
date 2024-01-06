## tools:
- NATS JetStream
- PostgreSQL
- Redis
- React

## setup:
```shell
cd configs
source .env
docker compose build
docker compose up
```
web: http://localhost:3000

#### script: 
```shell
go run publisher.go
```

## Vegeta
```shell
➜  ~ echo "GET http://localhost:8080/order/1" | vegeta attack -duration=30s | tee results.bin | vegeta report
Requests      [total, rate, throughput]         1500, 50.03, 50.03
Duration      [total, attack, wait]             29.984s, 29.98s, 4.085ms
Latencies     [min, mean, 50, 90, 95, 99, max]  702.375µs, 3.169ms, 3.064ms, 4.216ms, 5.016ms, 8.753ms, 21.824ms
Bytes In      [total, mean]                     1651500, 1101.00
Bytes Out     [total, mean]                     0, 0.00
Success       [ratio]                           100.00%
Status Codes  [code:count]                      200:1500
Error Set:
```
