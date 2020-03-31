# envoy-chaosproof-api

Statically configured envoy with active/passive health checking, primary/fallback cluster load balancing, circuit breakers, retries and timeouts.

```sh
make build

docker-compose up
```

## TODO

- Load test
- Add upstream auth layer / Lua script or JWT auth filter
- Synthetic testing / Check for testing user identity / Add headers / Lua script
- Add a canary subset https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/upstream/load_balancing/subsets
- TLS termination