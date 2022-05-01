<h1 align="center">Welcome </h1>

- This project is a simple API for basic operations of a Bank, such as creating accounts, get balance for a specific
  account, deposit money on a specific account, withdraw money from a specific account and get balance of a specific
  account.


## Architecture
I tried to design and implement this solution based on clean architecture (Ports and Adapters) standards and approaches.

## Web Framework
I used Gin Web Framework. Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter.

## Requirements/dependencies
- Docker
- Docker-compose

## Docker-compose up
if you want to see this example running, you can just type `docker-compose up` from solution directory.

## Docker-compose up -d
If you want to run this example but without attaching console, run _docker-compose up_ in detach mode - `docker-compose up -d`.

## Docker-compose up --build
If you have already composed system up, but then changed source code, you need to pass _--build_ parameter, when running _docker-compose up_ next time: `docker-compose up --build`.
Of course, it can be used along with detach parameter.

## Docker-compose down
When you want to clean up containers and networks created by _docker-compose_, just type `docker-compose down` from solution directory.

## Run tests in container
Please run this command on the Docker CLI command prompt:
```bash
go test -v ./...
```

## API Documentation

| Endpoint    | HTTP Method | Description       |
| ----------- |:-----------:| :-----------------: |
| `/accounts` |   `POST`    | `Create an account` |
| `/accounts/deposit` |    `PUT`    | `Deposit money on a specific account`   |
| `/accounts/withdraw` |    `PUT`    |    `Withdraw money from a specific account` |
| `/accounts/:id/balance`|    `GET`    | `Get balance of a specific account` |

## Test endpoints API using curl

- #### Create an account

`Request`

```bash
curl -i --request POST 'http://localhost:3001/accounts' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name":"ali",
    "last_name":"hamedani",
    "iban":"564564",
    "balance":11
}'
```

`Response`

```json
{
  "Status": true,
  "Result": {
    "id": 1514179864372969472,
    "first_name": "ali",
    "last_name": "hamedani",
    "iban": "564564",
    "balance": 11,
    "creation_date": "2022-04-13T11:53:06.4250305+02:00"
  },
  "Error": ""
}
```

- #### Deposit money

`Request`

```bash
curl -i --request PUT 'http://localhost:3001/accounts/deposit' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source_account_id": 1514179686534479872,
    "destination_account_id": 1514179864372969472,
    "amount":11
}'
```

`Response`

```json
{
  "Status": true,
  "Result": "success",
  "Error": ""
}
```

- #### Withdraw money

`Request`

```bash
curl -i --request PUT 'http://localhost:3001/accounts/withdraw' \
--header 'Content-Type: application/json' \
--data-raw '{
    "source_account_id": 1514179686534479872,
    "destination_account_id": 1514179864372969472,
    "amount":11
}'
```

`Response`

```json
{
  "Status": true,
  "Result": "success",
  "Error": ""
}
```

- #### Get balance

`Request`

```bash
curl -i --request GET 'http://localhost:3001/accounts/{id}/balance' 
```

`Response`

```json
{
  "Status": true,
  "Result": {
    "first_name": "ali",
    "last_name": "faghani",
    "iban": "45345353",
    "balance": 43
  },
  "Error": ""
}
```
## Test description
I have implemented domain model tests and use-case tests for this project.



## Security
I used simple token-based authentication in this project.
Every request should put the token value in the Authorization header.
The value of the token saved in API_TOKEN filed in the app.env configuration file.