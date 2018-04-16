Citadel Technologies - Budget Manager
=====================================

This project is a budget management tool using an API to store data about one or several budgets.

It stores transactions historic.

Install
-------

### Via Docker

First, pull the image.

```sh
docker pull citadeltechnologies/budget-manager
```

The container needs a MongoDB server running to store the data.

For the usage example, we will assume that you use the dedicated Docker container.

For simpler use, we create a network named ```budget``` to ease the connection between the two containers, but feel free to set up any kind of network configuration.

```sh
docker network create budget
docker run \
  -it \
  --name budget_mongo \
  --network budget \
  mongo
```

Then, launch the container with the proper environment values.

```sh
docker run \
  -it \
  -e "MONGO_HOST=budget_mongo" \
  -e "MONGO_PORT=27017" \
  -e "MONGO_DBNAME=test" \
  --network budget \
  -p 80:80 \
  citadeltechnologies/budget-manager
```

### Via Docker Compose

You can add the following service definition to your ```docker-compose.yml``` file.

```yml
budget-manager:
  container_name: ct_budget_manager
  image: citadeltechnologies/budget-manager
  environment:
    - MONGO_HOST=mongo
    - MONGO_PORT=27017
    - MONGO_DBNAME=test
  networks:
    - budget
  ports:
    - "80:80"
```

API
-----

### Routes

#### Transactions

GET ```/budgets```

Response 200
```json
[
  {
    "id": "56f4r86f5f45a6f62d",
    "name": "First year budget",
    "slug": "first-year-budget",
    "description": "Year 2015 budget",
    "sectors": [],
    "created_at": "2015-09-14T13:19:34.740462493Z",
    "updated_at": "2015-09-14T13:19:34.740462493Z"
  },
  {
    "id": "6regre4g56g435f12f",
    "name": "Previous year budget",
    "slug": "previous-year-budget",
    "description": "Year 2016 budget",
    "sectors": [],
    "created_at": "2016-09-14T13:19:34.740462493Z",
    "updated_at": "2016-09-14T13:19:34.740462493Z"
  },
  {
    "id": "5d5f4ds58gr4s5ds33",
    "name": "My current budget",
    "slug": "my-current-budget",
    "description": "Year 2017 budget",
    "sectors": [
        {
            "name": "Industry",
            "slug": "industry",
            "transactions": [
                {
                    "id": "5sdf4sfsdfds22dztra4a",
                    "wording": "New machine billing",
                    "description": "3 HB-265s were bought this month",
                    "type": "expense",
                    "amount": 15025.5,
                    "created_at": "2016-09-14T13:19:34.740462493Z",
                    "updated_at": "2016-09-14T13:19:34.740462493Z"
                }
            ]
        }
    ],
    "created_at": "2017-09-14T13:19:34.740462493Z",
    "updated_at": "2017-09-14T13:19:34.740462493Z"
  },
]

```
GET ```/budgets/{slug}```

Response 200
```json
{
    "id": "5d5f4ds58gr4s5ds33",
    "name": "My current budget",
    "slug": "my-current-budget",
    "description": "Year 2017 budget",
    "sectors": [
        {
            "name": "Industry",
            "slug": "industry",
            "transactions": [
                {
                    "id": "5sdf4sfsdfds22dztra4a",
                    "wording": "New machine billing",
                    "description": "3 HB-265s were bought this month",
                    "type": "expense",
                    "amount": 15025.5,
                    "created_at": "2016-09-14T13:19:34.740462493Z",
                    "updated_at": "2016-09-14T13:19:34.740462493Z"
                }
            ]
        }
    ],
    "created_at": "2017-09-14T13:19:34.740462493Z",
    "updated_at": "2017-09-14T13:19:34.740462493Z"
}
```

DELETE ```/budgets/{slug}```

Response 404
```json
{
    "code": 404,
    "message": "Budget not found"
}
```

Response 204

POST ```/budgets```

Request
```json
{
    "name": "My budget",
    "description": "Year 2017 budget"
}
```

Response 201
```json
{
    "id": "56f4r86f5f45a6f62d",
    "name": "My budget",
    "slug": "my-budget",
    "description": "Year 2017 budget",
    "sectors": [],
    "created_at": "2017-09-14T13:19:34.740462493Z",
    "updated_at": "2017-09-14T13:19:34.740462493Z"
}
```

POST ```/budgets/{budget-slug}/sectors```

Request
```json
{
    "name": "Industry"
}
```

Response 201
```json
{
    "name": "Industry",
    "slug": "industry",
    "transactions": []
}
```

GET ```/budgets/{budget-slug}/sectors/{sector-slug}/transactions```

Response 200
```json
[
  {
      "id": "59bba162a7bfdd0001db65e2",
      "wording": "Product 10230",
      "description": "Quantity - 2",
      "type": "income",
      "amount": 23.00,
      "created_at": "2017-09-15T09:46:10.098089142Z"
  },
  {
      "id": "59bba162a7bfdd0001db65e6",
      "wording": "Product 13542",
      "description": "Quantity - 1",
      "type": "income",
      "amount": 19.85,
      "created_at": "2017-09-12T12:53:41.098089142Z"
  },
  {
      "id": "59bba162a7bfdd0001db65ef",
      "wording": "Product 10230",
      "description": "Quantity - 2",
      "type": "income",
      "amount": 2396.00,
      "created_at": "2017-09-01T15:00:10.098089142Z"
  },
]
```

GET ```/budgets/{budget-slug}/sectors/{sector-slug}/transactions/{id}```

Response 200
```json
{
    "id": "59bba162a7bfdd0001db65b1",
    "wording": "Product 10230",
    "description": "Quantity - 2",
    "type": "income",
    "amount": 23.00,
    "created_at": "2017-09-15T09:46:10.098089142Z"
}
```

POST ```/budgets/{budget-slug}/sectors/{sector-slug}/transactions```

Request
```json
{
	"wording": "Product 10230",
	"description": "Quantity - 2",
	"type": "income",
	"amount": 23.00
}
```

Response 201
```json
{
    "id": "59bba162a7bfdd0001db65b1",
    "wording": "Product 10230",
    "description": "Quantity - 2",
    "type": "income",
    "amount": 23.00,
    "created_at": "2017-09-15T09:46:10.098089142Z"
}
```
