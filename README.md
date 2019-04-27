# Go-Graphql API Tutorial

## This a simple API created using golang w/ graphql

I'm following along with this [tutorial](https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356) by Bradford Lamson-Scribner to make this simple API.

## Getting Started

First I created a base template for the project. Following the structure that Bradford Lamson-Scribner set in his tutorial. I've made only a couple of minor changes to the structure.

### Project Structure

``` dir
├── gql
│   ├── gql.go
│   ├── queries.go
│   ├── resolvers.go
│   └── types.go
├── main.go
├── db
│   └── postgres.go
└── server
    └── server.go
```

## Setting up a Postgres DB in Docker

### Settings up a docker instance of postgres

I found a [great article](https://hackernoon.com/dont-install-postgres-docker-pull-postgres-bee20e200198) by Syed Komail Abbas that goes into detail to get started with a docker container with postgres db running in it.

#### Download the latest Postgres Docker Image

To pull the latest docker image of postgres simply run

``` zsh
docker pull postgres
```

#### Create a Directory to Serve as the Local Host Point for the Postgres Container's Data Files

If you want persistant data generated by the Postgres instance, we need to map a local directory as a data volume for the container.

``` zsh
mkdir -p $HOME/docker/volumes/postgres
```

#### Run the Postgres Container

To start the postgres container simply run

``` zsh
docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 -v $HOME/docker/volumes/postgres:/var/lib/postgresql/data postgres
```

#### Connecting to the Postgres Container

Once the container is up and running, connecting to it from an application is no different than connecting to a Postgres instance running outside a docker container.

``` zsh
psql -h localhost -U postgres -d postgres
```

### Creating a DB for the API

We need to setup a database and some mocked data to test the API against. Run ```psql``` to enter the Postgres console and we'll create a database.

After connecting to the Postgres container run

``` sql
CREATE DATABASE go_graphql_db;
```

Then we'll connect to the new database that we created.

``` sql
\c go_graphql_db;
```

Once the database is created, go ahead and populate the database with some mock data.

``` sql
CREATE TABLE users (
  id serial PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  age INT NOT NULL,
  profession VARCHAR (50) NOT NULL,
  friendly BOOLEAN NOT NULL
);

INSERT INTO users VALUES
  (1, 'kevin', 35, 'waiter', true),
  (2, 'angela', 21, 'concierge', true),
  (3, 'alex', 26, 'zoo keeper', false),
  (4, 'becky', 67, 'retired', false),
  (5, 'kevin', 15, 'in school', true),
  (6, 'frankie', 45, 'teller', true);
```

## API

See this [tutorial](https://medium.com/@bradford_hamilton/building-an-api-with-graphql-and-go-9350df5c9356) to get the code for the api.

## Testing the API

There are some great tools for exploring your GraphQL API like [graphiql](https://github.com/graphql/graphiql), [Insomnia](https://insomnia.rest/), and [graphql-playground](https://github.com/prisma/graphql-playground). You can also just make a POST request sending over a raw application/json body like this:

``` json
{
  "query": "{users(name:\"kevin\"){id, name, age}}"
}
```