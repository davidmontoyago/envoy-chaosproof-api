# envoy-chaosproof-api

Statically configured envoy with active/passive health checking, primary/fallback cluster load balancing, circuit breakers, retries and timeouts.

```sh
make build

docker-compose up
```

## TODO

- Fallbacks? https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http_routing#direct-responses
- Buffering
- Add a canary subset https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/subsets
- TLS termination