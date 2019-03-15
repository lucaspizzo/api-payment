# api-payment

### Para rodar a aplicação basta rodar :
```shell 
docker-compose up
```


---

### **PATCH** `/v1/accounts/:id`
#### Request

```json
{
	"available_credit_limit": {
	  "amount": 123.45
	},
	"available_withdraw_limit": {
      "amount": 123.45
    }
}
```
#### Response
**Status**
    `200`
```json
{
  "available_credit_limit": {
      "amount": 200.45
  },
  "available_withdraw_limit": {
      "amount": 200.45
  }
}
```
---
### **GET** `/v1/accounts/limits`
#### Response
**Status**
    `200`
```json
[
    {
      "available_credit_limit": {
          "amount": 200.45
      },
      "available_withdraw_limit": {
          "amount": 200.45
      }
    }
]

```
---

### **POST** `/v1/transactions`
#### Request

```json
{
	"account_id": 1,
	"amount": 200.2,
	"operation_type_id": 1
}
```
#### Response
**Status**
    `200`
```json
{
  "account_id": 1,
  "operation_type_id": 1,
  "amount": 200.2,
  "balance": -200.2,
  "event_date": "20/02/2002",
  "due_date": "20/02/2002"
}
```
---

### **POST** `/v1/payments`
#### Request

```json
[
  {
	"account_id": 1,
	"amount": 200.2
  },
  {
  	"account_id": 1,
  	"amount": 25
  }
]
```
#### Response
**Status**
    `204`
