# LINK-AJA API

Simple REST API 

## Installation

```bash
docker-compose up
```

### Step
1. Open phpmyadmin on your browser (http://localhot:9090)
2. Create Database `link-aja`
3. Import database `link-aja.sql`
4. Curl for Check Saldo
    ```bash
    curl --location --request GET 'http://localhost:8080/account/555001'
    ```
5. Curl for Transfer
    ```bash
    curl --location --request POST 'http://localhost:8080/account/555001/transfer' \
    --header 'Content-Type: application/json' \
    --data-raw '{
        "to_account_number": "555002",
        "amount": 100
    }'
    ```