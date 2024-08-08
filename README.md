# Take home assignment

## Description:

There's 2 services for handle each domain:

1. **go-payment-srv**: Golang based service to handle transaction processes like withdraw and send money.
2. **node-account-srv**: Node based service to handle account creation and authentication

For details, you can check here: https://www.notion.so/Concrete-ai-Test-Docs-9d4197ab0d4b4d62ac4a2cdfd91716d3?pvs=4

## How To Run

1. run these command

```console
docker-compose -f deployment/docker-compose.yml up -d --build
```

2. after container success deployed, go to `services/node-account-srv` then run `./migrate.sh` on your local machine
3. check two services is ready to serve
   1. go-payment-srv: `http://localhost:8001`
   2. node-account-srv: `http://localhost:8000`
