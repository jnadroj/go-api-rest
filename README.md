## Creating the database

Let's create the database which is a container of mongodb

### Getting started

First we have to run the command to raise the docker service:

```bash
docker-compose up -d
```

Now we must enter the container with the command:

```bash
docker exec -it mongodb bash
```

Once we are inside the container we can log into mongo with:

```bash
mongosh -u root -p root
```

Now we are inside mongodb.
