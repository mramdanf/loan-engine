### Run locally
- run server `go run main.go`
- run unit test `go test ./...`

### Create Customer Loan
```
curl --location 'http://localhost:8080/customer-loan' \
--header 'Content-Type: application/json' \
--data '{
    "customer_id": 1,
    "loan_id": 1
}'
```

### Make payment
Assume can target specific billing date. Real implementation should use current date from system
```
curl --location 'http://localhost:8080/customer-loan/payment' \
--header 'Content-Type: application/json' \
--data '{
    "customer_loan_id": 0,
    "payment_date": "2025-05-22 21:42:06"
}'
```

### Get loan outstanding
```
curl --location 'http://localhost:8080/customer-loan/0/outstanding'
```

### Check delinquent
Assume can target specific billing date. Real implementation should use current date from system
```
curl --location 'http://localhost:8080/customer-loan/0/delinquent?current_date=2025-06-05%2021%3A42%3A06'
```