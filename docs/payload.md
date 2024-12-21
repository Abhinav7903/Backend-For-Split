Here’s a breakdown of the expected JSON formats for the various database operations (such as creating, retrieving, updating, deleting, and searching transactions) in your PostgreSQL operations.

---

### 1. **Create Transaction** (`POST /create-transaction`)

#### Request Body (JSON):

```json
{
  "lender_id": 1,
  "borrower_id": 2,
  "group_id": 1,
  "amount": 500.00,
  "status": "pending",
  "purpose": "Loan repayment",
  "payment_method_id": 1,
  "retry_count": 0,
  "failure_reason": null
}
```

- `lender_id`: ID of the lender.
- `borrower_id`: ID of the borrower.
- `group_id`: ID of the group (nullable).
- `amount`: Amount for the transaction.
- `status`: Status of the transaction (e.g., `"pending"`, `"completed"`).
- `purpose`: Purpose of the transaction (nullable).
- `payment_method_id`: ID of the payment method (nullable).
- `retry_count`: The number of retries (integer).
- `failure_reason`: Reason for failure (nullable).

#### Response Body (JSON):

```json
{
  "message": "Transaction created",
  "data": 12345
}
```

- `message`: Success message.
- `data`: The `transaction_id` of the created transaction.

---

### 2. **Get Transaction by ID** (`GET /get-transaction?id=<transaction_id>`)

#### Request (URL Parameter):
```
GET /get-transaction?id=12345
```

#### Response Body (JSON):

```json
{
  "message": "Transaction retrieved",
  "data": {
    "transaction_id": 12345,
    "lender_id": 1,
    "borrower_id": 2,
    "group_id": 1,
    "amount": 500.00,
    "status": "pending",
    "purpose": "Loan repayment",
    "payment_method_id": 1,
    "retry_count": 0,
    "failure_reason": null
  }
}
```

- `message`: Success message.
- `data`: The transaction details.

---

### 3. **Get Transactions by Lender ID** (`GET /get-transactions-by-lender?lender_id=<lender_id>`)

#### Request (URL Parameter):
```
GET /get-transactions-by-lender?lender_id=1
```

#### Response Body (JSON):

```json
{
  "message": "Transactions retrieved",
  "data": [
    {
      "transaction_id": 12345,
      "lender_id": 1,
      "borrower_id": 2,
      "group_id": 1,
      "amount": 500.00,
      "status": "pending",
      "purpose": "Loan repayment",
      "payment_method_id": 1,
      "retry_count": 0,
      "failure_reason": null
    },
    {
      "transaction_id": 12346,
      "lender_id": 1,
      "borrower_id": 3,
      "group_id": 1,
      "amount": 1000.00,
      "status": "completed",
      "purpose": "Loan repayment",
      "payment_method_id": 2,
      "retry_count": 0,
      "failure_reason": null
    }
  ]
}
```

- `message`: Success message.
- `data`: An array of transaction objects.

---

### 4. **Get Transactions by Borrower ID** (`GET /get-transactions-by-borrower?borrower_id=<borrower_id>`)

#### Request (URL Parameter):
```
GET /get-transactions-by-borrower?borrower_id=2
```

#### Response Body (JSON):

```json
{
  "message": "Transactions retrieved",
  "data": [
    {
      "transaction_id": 12345,
      "lender_id": 1,
      "borrower_id": 2,
      "group_id": 1,
      "amount": 500.00,
      "status": "pending",
      "purpose": "Loan repayment",
      "payment_method_id": 1,
      "retry_count": 0,
      "failure_reason": null
    }
  ]
}
```

- `message`: Success message.
- `data`: An array of transaction objects.

---

### 5. **Update Transaction Status** (`PUT /update-transaction-status?id=<transaction_id>`)

#### Request Body (JSON):

```json
{
  "status": "completed"
}
```

- `status`: The updated status for the transaction (e.g., `"completed"`, `"failed"`).

#### Response Body (JSON):

```json
{
  "message": "Transaction status updated"
}
```

- `message`: A success message confirming the status update.

---

### 6. **Delete Transaction** (`DELETE /delete-transaction?id=<transaction_id>`)

#### Request (URL Parameter):
```
DELETE /delete-transaction?id=12345
```

#### Response Body (JSON):

```json
{
  "message": "Transaction deleted"
}
```

- `message`: Success message indicating the transaction was deleted.

---

### 7. **Search Transactions** (`POST /search-transactions`)

#### Request Body (JSON):

```json
{
  "lender_id": 1,
  "borrower_id": 2,
  "group_id": 1,
  "status": "pending",
  "min_amount": 100.00,
  "max_amount": 1000.00,
  "payment_method_id": 1
}
```

- `lender_id`: ID of the lender (optional filter).
- `borrower_id`: ID of the borrower (optional filter).
- `group_id`: ID of the group (optional filter).
- `status`: Transaction status (optional filter).
- `min_amount`: Minimum transaction amount (optional filter).
- `max_amount`: Maximum transaction amount (optional filter).
- `payment_method_id`: ID of the payment method (optional filter).

#### Response Body (JSON):

```json
{
  "message": "Transactions retrieved",
  "data": [
    {
      "transaction_id": 12345,
      "lender_id": 1,
      "borrower_id": 2,
      "group_id": 1,
      "amount": 500.00,
      "status": "pending",
      "purpose": "Loan repayment",
      "payment_method_id": 1,
      "retry_count": 0,
      "failure_reason": null
    },
    {
      "transaction_id": 12346,
      "lender_id": 1,
      "borrower_id": 2,
      "group_id": 1,
      "amount": 800.00,
      "status": "completed",
      "purpose": "Loan repayment",
      "payment_method_id": 2,
      "retry_count": 1,
      "failure_reason": null
    }
  ]
}
```

- `message`: Success message.
- `data`: An array of transactions that match the search filters.

---

### Summary of JSON Fields:

- **Create Transaction**: Fields to provide details about the transaction.
- **Get Transaction**: Returns transaction details for a specific ID.
- **Get Transactions by Lender/Borrower ID**: Returns an array of transactions for a given lender or borrower.
- **Update Transaction Status**: Accepts a new status for a transaction.
- **Delete Transaction**: Confirms that a transaction was deleted.
- **Search Transactions**: Allows filtering transactions by multiple criteria.
