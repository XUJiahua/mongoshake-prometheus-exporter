# mongoshake-prometheus-exporter

Convert MongoShake metrics (http://localhost:9100/repl) to Prometheus metrics.

## local test

```
go run main.go testdata/config.toml
```

## Helm chart

[mongoshake-prometheus-exporter](mongoshake-prometheus-exporter)

## reference
1. repl https://github.com/alibaba/MongoShake/wiki/%E5%A6%82%E4%BD%95%E7%9B%91%E6%8E%A7%E5%92%8C%E7%AE%A1%E7%90%86MongoShake%E7%9A%84%E8%BF%90%E8%A1%8C%E7%8A%B6%E6%80%81%EF%BC%9F#21-repl
2. Q: How to monitor the MongoShake? https://github.com/alibaba/MongoShake/wiki/FAQ#q-how-to-monitor-the-mongoshake
