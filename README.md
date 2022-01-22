golang-fiber-tutorial-implementation

- Creating the network that the containers will communicate
  `docker network create -d bridge dev-network`

- Creating the postgresql container

```
docker run --rm -d \
    --name dev-postgres \
    --network dev-network \
    -e POSTGRES_USER=postgres \
    -e POSTGRES_PASSWORD=password \
    -e POSTGRES_DB=postgres \
    -v ${HOME}/dev-postgres/data/:/var/lib/postgresql/data \
    -p 5432:5432 \
    postgres:14
```

