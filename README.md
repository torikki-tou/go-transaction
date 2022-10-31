# go-transaction

### How to run

`docker-compose up --build`

### How to use

`0.0.0.0:5050` - PDAdmin

- User: admin@admin.com
- Password: root

Add new DB:

- Host: db
- User: postgres

`0.0.0.0:8000/api/client/change_balance` - balance changing endpoint

Request examples:

```
{
    "client_id": "aAqlNVt",
    "delta": 10
}
```

```
{
    "client_id": "aAqlNVt",
    "delta": -10
}
```

`0.0.0.0:15672` - RabbitMQ GIU

- User: guest
- Password: guest
