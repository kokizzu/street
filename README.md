
# Street Project

People can buy/sell house/building

- main: inject all dependencies

- presentation -> serialization and transport
- domain -> business logic, DTO
- model -> persistence/3rd party endpoints, DAO

## CQRS

- rq = request/query (OLTP)
- wc = write/command (OLTP)
- sa = statistics/analytics (OLAP)

## start dev mode

```shell
# start go server
air 

# manually
go run main.go web

# start frontend auto build 
cd svelte
npm run watch
```

## Manual test

```shell
# using command line
go run main.go cli UserRegister '{"email":"a@b.c"}'

# using curl
go run main.go web
curl -X POST -d '{"email":"test@a.com"}' localhost:1234/UserRegister
```