# Kafka Connect MySQL Sink Example

An example of using [Kafka Connect](https://docs.confluent.io/platform/current/connect/index.html) with the JDBC sink connector to copy messages from a Kafka topic into a MySQL table. A Go producer writes 10,000 JSON messages, and Kafka Connect upserts them into MySQL with no consumer code of your own.

```
main.go ──▶ Kafka topic say-hello ──▶ Kafka Connect (JDBC sink) ──▶ MySQL my_db.say_hello
```

## Pieces

- [docker-compose.yml](docker-compose.yml): ZooKeeper and Kafka from Confluent Platform 7.9, MySQL 8.4, and Kafka Connect in standalone mode. The Connect container installs the JDBC connector (10.9.9) and MySQL Connector/J (9.7.0) when it starts.
- [init-db.sql](init-db.sql): creates the `say_hello` table (`id`, `name`, `created_at`, `updated_at`).
- [connect-standalone.properties](connect-standalone.properties): the worker config. Values use the `JsonConverter` with schemas enabled.
- [connect-mysql-sink.properties](connect-mysql-sink.properties): the sink config, which reads topic `say-hello`, writes table `say_hello` in `upsert` mode, and uses the `id` field of the value as the primary key.
- [main.go](main.go): sends messages with IDs 1 to 10,000 to `say-hello`. Each message carries the schema from [schema.json](schema.json) next to its payload, because the JDBC sink needs a schema to map fields to columns:

```json
{
  "schema": { "type": "struct", "fields": [{ "type": "int64", "field": "id" }, { "type": "string", "field": "name" }] },
  "payload": { "id": 1, "name": "Rashad 1" }
}
```

## Run

Requires Go and Docker.

```bash
make up         # start ZooKeeper, Kafka and MySQL
make connect    # start Kafka Connect and follow its log until the connector is running
make generate   # in another terminal: produce 10,000 messages
```

Check the result:

```bash
docker exec -it mysql mysql -umy_user -pmy_pass my_db -e 'SELECT COUNT(*) FROM say_hello'
```

Running `make generate` again updates the same 10,000 rows instead of adding new ones, because the sink upserts by `id`.

```bash
make down
```

## References

- [Kafka Connect quick start](https://docs.confluent.io/3.2.2/connect/quickstart.html)
- [Installing Kafka Connect plugins](https://docs.confluent.io/home/connect/install.html)
- [JDBC connector](https://docs.confluent.io/kafka-connect-jdbc/current/index.html)
- [Kafka Connect converters and serialization](https://www.confluent.io/blog/kafka-connect-deep-dive-converters-serialization-explained/)

## License

[MIT](LICENSE)
