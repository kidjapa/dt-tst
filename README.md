# DT-TST

```
Go: 1.16.5
```

Environment Variables
---------------------
Name | Description | Accept Values | Default |
-----|-----------|---------|--------------|
ENV_PORT | Define server port | ```9999``` | ```9090```

Annotations
-----------
- How to use .tool-version: [Here](https://github.com/asdf-vm/asdf)
- Run docker: ```docker-compose -f ./docker-compose.yml up -d```

Determined tests Passed
-----------------------
- [x] Reset state before starting tests
- [x] Get balance for non-existing account
- [x] Create account with initial balance
- [x] Deposit into existing account
- [x] Get balance for existing account
- [x] Withdraw from non-existing account
- [x] Withdraw from existing account
- [x] Transfer from existing account
- [x] Transfer from non-existing account

```text
--
# Reset state before starting tests

POST /reset

200 OK


--
# Get balance for non-existing account

GET /balance?account_id=1234

404 0


--
# Create account with initial balance

POST /event {"type":"deposit", "destination":"100", "amount":10}

201 {"destination": {"id":"100", "balance":10}}


--
# Deposit into existing account

POST /event {"type":"deposit", "destination":"100", "amount":10}

201 {"destination": {"id":"100", "balance":20}}


--
# Get balance for existing account

GET /balance?account_id=100

200 20

--
# Withdraw from non-existing account

POST /event {"type":"withdraw", "origin":"200", "amount":10}

404 0

--
# Withdraw from existing account

POST /event {"type":"withdraw", "origin":"100", "amount":5}

201 {"origin": {"id":"100", "balance":15}}

--
# Transfer from existing account

POST /event {"type":"transfer", "origin":"100", "amount":15, "destination":"300"}

201 {"origin": {"id":"100", "balance":0}, "destination": {"id":"300", "balance":15}}

--
# Transfer from non-existing account

POST /event {"type":"transfer", "origin":"200", "amount":15, "destination":"300"}

404 0
```
