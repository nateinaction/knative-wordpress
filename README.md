# knative-wordpress

This repository contains a Knative service for running WordPress with PHP and Nginx.

## Kind Setup

Spin up a kind cluster with Knative installed:

```sh
# install kind
brew install kind

# install knative cli
brew install knative/client/kn

# install knative quickstart plugin
brew install knative-extensions/kn-plugins/quickstart

# create a kind cluster with knative
kn quickstart kind
```

Now deploy the application
```sh
kubectl apply -k config
```

## Database Setup

Create the database

```sh
# Get the root password from the secret
kubectl get secret -n db internal-cluster1 --template='{{.data.root | base64decode}}{{"\n"}}'

# Create a shell in a MySQL client pod
kubectl run --rm -i -n db --tty mysql-client --image=mysql:latest --restart=Never -- bash -il

# Connect to the MySQL server
mysql -h cluster1-haproxy -u root -p '<root_password>'

# Create the WordPress database
CREATE DATABASE wordpress;

# Create a user for WordPress
CREATE USER 'wordpress'@'%' IDENTIFIED BY '<password>';
GRANT ALL PRIVILEGES ON wordpress.* TO 'wordpress'@'%';
FLUSH PRIVILEGES;
```

Finally, set the database password in the `config/wordpress/secret.yaml` file.

## Load Testing

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
