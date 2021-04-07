# Transactions #

### Prerequisites

You need `Docker` and `docker-compose`. Install docker compose [documentation](https://docs.docker.com/compose/install/)

### Principal Commands

- Run application with docker on port `:8080`
```shell
make up
```

- Show and attach to application logs
```shell
make logs
```

- Run tests with coverage
```shell
make test
```

## API Endpoints

| Endpoint | Method | Description |
| :----------------: | :-------------------: | :-------------------: |
| `/accounts` | `POST` | `Create Account` |
| `/accounts/{:accountId}` | `GET` | `Get account by id` |
| `/transactions` | `POST` | `Create Transaction` |
| `/transactions/{:transactionId}` | `GET` | `Get transaction by id` |



### References
- https://eminetto.medium.com/clean-architecture-using-golang-b63587aa5e3f
- https://eltonminetto.dev/en/post/2020-07-06-clean-architecture-2years-later/
- https://hackernoon.com/golang-clean-archithecture-efd6d7c43047
