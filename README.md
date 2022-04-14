<h1 align="center">Welcome </h1>

- This project is a simple API for basic operations of a Bank, such as creating accounts, get balance for a specific
  account, deposit money on a specific account, withdraw money from a specific account and get balance of a specific
  account.


## Architecture
I tried to design and implement this solution based on clean architecture (Ports and Adapters) standards and approaches.

## Web Framework
I used Gin Web Framework. Gin is a web framework written in Go (Golang). It features a martini-like API with performance that is up to 40 times faster thanks to httprouter.
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