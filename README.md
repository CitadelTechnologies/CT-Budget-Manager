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
  -e "MONGO_DBNAME=test"
  -e "SERVER_PORT=80" \
  --network budget
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
    - SERVER_PORT=80
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

Response
```json
[
  {
    "id": "56f4r86f5f45a6f62d",
    "name": "First year budget",
    "description": "Year 2015 budget",
    "transactions": [],
    "created_at": "2015-09-14T13:19:34.740462493Z",
    "updated_at": "2015-09-14T13:19:34.740462493Z"
  },
  {
    "id": "6regre4g56g435f12f",
    "name": "Previous year budget",
    "description": "Year 2016 budget",
    "transactions": [],
    "created_at": "2016-09-14T13:19:34.740462493Z",
    "updated_at": "2016-09-14T13:19:34.740462493Z"
  },
  {
    "id": "5d5f4ds58gr4s5ds33",
    "name": "My current budget",
    "description": "Year 2017 budget",
    "transactions": [],
    "created_at": "2017-09-14T13:19:34.740462493Z",
    "updated_at": "2017-09-14T13:19:34.740462493Z"
  },
]
```

POST ```/budgets```

Request
```json
{
  "name": "My budget",
  "description": "Year 2017 budget"
}
```

Response
```json
{
  "id": "56f4r86f5f45a6f62d",
  "name": "My budget",
  "description": "Year 2017 budget",
  "transactions": [],
  "created_at": "2017-09-14T13:19:34.740462493Z",
  "updated_at": "2017-09-14T13:19:34.740462493Z"
}
```

GET ```/transactions```

GET ```/transaction```

POST ```/transaction```
