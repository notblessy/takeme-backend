# takeme-app Backend Service

takemeApp is a simple backend service for dating app. The name takemeApp is a pun for telling "take me up".

## Features

- Register & Login
- Find Users
- Create Reaction (swipe left or right) with caching
- Premium features to unlock limitations
- Run in Docker

## Architecture

- takeme-app uses a clean architecture referenced by go-clean-arch `https://github.com/bxcodec/go-clean-arch`. It has 4 domain layers such

  - Models Layer
  - Repository Layer
  - Usecase Layer
  - Delivery Layer

- Here's some reason why referencing go-clean-arch:
  - Reliable:
    - Domains are easly and independently testable.
    - Adapting Domainn driven design & SOLID principles.
  - Maintainable
    - A way more organized code. Even with separated layers, still encapsulates logic business that may not affect other logic when having updates/changes.
    - Code should not depend on one developer. Readable codes will saves time and effort of developers in future.

## Stacks

- golang
- echo HTTP Framework
- .modd
- dbmate
- mysql
- redis
- makefile `brew install make`

## Clone

First clone this repo by run:

```sh
$ git clone git@github.com:notblessy/takeme-backend.git
```

## Run takeme-app in local

### Initialize

Firstly, run:

```sh
$ go mod tidy
```

### Environment

- The sample environments are provided in root folder
  - If you run takeme-app in local, use `config.yml.sample` to be `config.yml` file.

### Database Migration

- Ensure you have already installed `Makefile` and created the database. To migrate tables, run:

```sh
$ make migration
```

- To seed necessary data:

```sh
$ make seeder
```

### Running app

- To run HTTP server with hot reloading by listens changes, hit:

```sh
$ make run
```

## API Documentations

- To test the API, import `postman collection` from folder `api-docs/`. All the API is available there.
- Create environtment in postman, there are 2 necessary variables for testing.
  - host. represents basehost value
  - token. represents auth token

#

## Author

```
I Komang Frederich Blessy
```
