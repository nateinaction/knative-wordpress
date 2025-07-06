# Database Cache Notes

An analysis of Redis based database cache options for serverless WordPress sites running on Knative.

## Predis
According to the [redis-cache](https://github.com/rhubarbgroup/redis-cache/?tab=readme-ov-file#configuration) WordPress plugin documentation, leader/follower replication is only available with the Predis PHP extension.

> Predis is the recommended PHP client for Redis.
> https://redis.io/docs/latest/develop/clients/php/

## Relay
Alternatively, Relay could be used instead of having an in-service redis replica. Relay would have colder starts compared to synced replicas but may scale well if tuned.

> PhpRedis and Relay perform significantly better when network I/O is involved, due to its ability to compress data by ~75%. Fewer bytes and received sent over the network means faster operations, and potentially cost savings when network traffic isn't free (e.g. AWS ElastiCache Inter-AZ transfer costs).
> https://github.com/predis/predis/blob/main/FAQ.md#when-should-i-use-phpredis

## PhpRedis
The PhpRedis extension does not offer replication support and does not seem suited for the serverless use case.

## SQLite
Another alternative is to forgo the use of a live database altogether and instead to use a flavor of sqlite since it appears to be a [beta datastore](https://github.com/WordPress/sqlite-database-integration) for WordPress. Something like [rqlite](https://github.com/rqlite/rqlite) feels about right given it's focus on distributed scaling.

### Options
- Predis with one redis leader and a read replica in each serverless workload
- Relay php extension installed with one redis leader and no read replicas
- Sqlite with no live database or db cache in the serverless workloads
