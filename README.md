# knative-wordpress

This repository contains a Knative service for running WordPress with PHP and Nginx.

```sh
kubectl apply -f service.yaml
```

Use `hey` for load testing:

```sh
go install github.com/rakyll/hey@latest
```

```sh
hey -z 30s -c 50 "http://wordpress.default.127.0.0.1.sslip.io"
```

Example output:
```sh
Summary:
  Total:	30.2637 secs
  Slowest:	3.3233 secs
  Fastest:	0.1471 secs
  Average:	0.3875 secs
  Requests/sec:	128.5370
  

Response time histogram:
  0.147 [1]	|
  0.465 [3225]	|■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■
  0.782 [449]	|■■■■■■
  1.100 [95]	|■
  1.418 [23]	|
  1.735 [48]	|■
  2.053 [25]	|
  2.370 [17]	|
  2.688 [3]	|
  3.006 [0]	|
  3.323 [4]	|


Latency distribution:
  10% in 0.1984 secs
  25% in 0.2491 secs
  50% in 0.3163 secs
  75% in 0.4093 secs
  90% in 0.5688 secs
  95% in 0.8093 secs
  99% in 1.8880 secs

Details (average, fastest, slowest):
  DNS+dialup:	0.0001 secs, 0.1471 secs, 3.3233 secs
  DNS-lookup:	0.0001 secs, 0.0000 secs, 0.0070 secs
  req write:	0.0000 secs, 0.0000 secs, 0.0001 secs
  resp wait:	0.3089 secs, 0.1321 secs, 3.2556 secs
  resp read:	0.0028 secs, 0.0000 secs, 0.0958 secs

Status code distribution:
  [200]	3890 responses
```
