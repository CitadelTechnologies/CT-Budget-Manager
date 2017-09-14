Citadel Technologies - Budget Manager
=====================================

This project is a budget management tool using an API to store data about one or several budgets.

It stores transactions historic.

Install
-------

### Via Docker

First, pull the image.

```sh
docker pull citadel-technologies/budget-manager
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
  -e "SERVER_PORT=80" \
  --network budget
  -p 80:80 \
  citadel-technologies/budget-manager
```

### Via Docker Compose

You can add the following service definition to your ```docker-compose.yml``` file.

```yml
budget-manager:
  container_name: ct_budget_manager
  image: citadel-technologies/budget-manager
  environment:
    - MONGO_HOST=mongo
    - MONGO_PORT=27017
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

GET ```/transactions```

GET ```/transaction```

POST ```/transaction```
