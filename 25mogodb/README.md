there is a special website for drivers using with golang

go.mongodb.org

we will need these golang library for mongodb
```
go get go.mongodb.org/mongo-driver/mongo
```

use a docker image for mongodb

compose file
```
version: '3.8'

services:
  mongo:
    image: mongo:latest
    container_name: some-mongo
    environment:
      MONGO_INITDB_ROOT_USERNAME: mongoadmin
      MONGO_INITDB_ROOT_PASSWORD: secret
    volumes:
      - ./mongo-data:/data/db
    ports:
      - "27017:27017"
    restart: unless-stopped
```

start the container
```
docker compose up -d
```

Use the tool MongoDB compass to test a connection

