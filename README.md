## Execution Flow

### Area Execution Workflow

* When an Area is triggered:
    1. Validate the Area’s configuration (ensure all Actions/Reactions are valid).
    2. Execute Actions in parallel or sequentially based on their `IsSequential` flag.
    3. Pass the results of Actions to Reactions.
    4. Trigger Reactions in parallel or sequentially as specified.

### Flow examples

1. **Simple Flow**:  
   `Action (Webhook triggered)` -> `Reaction (Send Email)`
2. **Parallel Actions**:  
   `Action 1 (File Uploaded)` + `Action 2 (API Call)` -> `Reaction (Send Notification)`
3. **Chained Workflow**:  
   `Action (Webhook)` -> `Reaction (API Call)` -> `Action (API Response)` -> `Reaction (Send Slack Message)`

## Kong Gateway Setup

### Installation (Docker)

```shell
docker network create kong-net
```

```shell
docker run -d --name kong-database \
     --network=kong-net \
     -p 5432:5432 \
     -e "POSTGRES_USER=kong" \
     -e "POSTGRES_DB=kong" \
     -e "POSTGRES_PASSWORD=kongpass" \
     postgres:13
```

```shell
docker run --rm --network=kong-net \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_PASSWORD=kongpass" \
    kong/kong-gateway kong migrations bootstrap
```

```shell
docker run -d --name kong-gateway \
    --network=kong-net \
    -e "KONG_DATABASE=postgres" \
    -e "KONG_PG_HOST=kong-database" \
    -e "KONG_PG_USER=kong" \
    -e "KONG_PG_PASSWORD=kongpass" \
    -e "KONG_PROXY_ACCESS_LOG=/dev/stdout" \
    -e "KONG_ADMIN_ACCESS_LOG=/dev/stdout" \
    -e "KONG_PROXY_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_ERROR_LOG=/dev/stderr" \
    -e "KONG_ADMIN_LISTEN=0.0.0.0:8001, 0.0.0.0:8444 ssl" \
    -e "KONG_ADMIN_GUI_URL=http://localhost:8002" \
    -p 8000:8000 \
    -p 8443:8443 \
    -p 127.0.0.1:8001:8001 \
    -p 127.0.0.1:8002:8002 \
    -p 127.0.0.1:8444:8444 \
    kong/kong-gateway
```

### Verify Installation

```shell
curl -i -X GET --url http://localhost:8001/services
```

### Admin panel

```shell
http://localhost:8002
```

### Create a Service

```shell
curl -i -X POST --url http://localhost:8001/services/ \
  --data 'name=example-service' \
  --data 'url=http://httpbin.org'
```

### Cleanup

```shell
docker kill kong-gateway
docker kill kong-database
docker container rm kong-gateway
docker container rm kong-database
docker network rm kong-net
```