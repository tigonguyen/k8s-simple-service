# simple-service

`simple-service` is an 'as easy as possible' golang HTTP api used at Onefootball on examples and tests.

It has only a `/live` endpoint answering `text/plain; charset=utf-8`. The following responses are possibly:
- `Well done :)`: if the application was able to connect with a Postgres database
- `Running`: if some error occurred during the connection with the database

Check the [`config`](/config/) package for more details on how to configure the database connection.

## Building

`simple-service` uses [go modules](https://github.com/golang/go/wiki/Modules). To build the application you can use the following command:

```
GO111MODULE=on go build
```

The command above will generate a binary `simple-service` that starts the application.
