
# Street Project

People can buy/sell house/building

[links](https://docs.google.com/document/d/1ATjUVawcWBM_axQBJU2Xdcu2WkF4sz9-NSGCSDzBHJ8/edit)

## Architecture

nothing fancy, just normal monolith but with multiple database connection

```
browser --> svelte+monolith --> tarantool/clickhouse
mobile  --> monolith --> tarantool/clickhouse
```

## Project Structure

- `conf`: shared configuration
- `main`: inject all dependencies
- `svelte`: web frontend
- `deploy`: scripts for deploying to production

MVC-like structure

- `presentation` -> serialization and transport
- `domain` -> business logic, DTO
- `model` -> persistence/3rd party endpoints, DAO

## CQRS

separates read/query and write/command into different database/connection.

- `rq` = read/query (OLTP) -> to tarantool replica
- `wc` = write/command (OLTP) -> to tarantool master
- `sa` = statistics/analytics (OLAP) -> to clickhouse

## Start dev mode

```shell
# start golang backend server
air web

# manually
go run main.go web

# start frontend auto build 
cd svelte
npm run watch
```

## Generate ORM

```shell
# input: 
# - model/m*.go

./gen-orm.sh
# or
cd model
go test -bench=BenchmarkGenerateOrm
# then go generate each file

# output:
# - model/m*/rq*/*.go  # -- read/query models
# - model/m*/wc*/*.go  # -- write/command mutation models
# - model/m*/sa*/*.go  # -- statistic/analytics models
```

## Generate Views

```shell
# input: 
# - domain/*.go
# - model/m*/*/*.go
# - svelte/*.svelte

cd presentation
go get -bench=BenchmarkGenerateViews

# output:
# - presentation/actions.GEN.go     # -- all possible commands
# - presentation/api_routes.GEN.go  # -- automatic API routes
# - presentation/web_view.GEN.go    # -- all template that can be used in web_static.go
# - presentation/cmd_run.GEN.go     # -- all CLI commands
# - svelte/jsApi.GEN.js             # -- all API client SDK 
```

## Manual test

```shell
# using command line
go run main.go cli guest/register '{"email":"a@b.c"}'

# using curl
go run main.go web

curl -X POST -d '{"email":"test@a.com"}' localhost:1234/guest/register
```

## Note

```shell
# docker spawning failed (because test terminated improperly), run this:
alias dockill='docker kill $(docker ps -q); docker container prune -f; docker network prune -f'
```

## FAQ

- **Q**: where to put SSR?
  - **A**: `presentation/web_static.go`
- **Q**: got error `there is no space with name [tableName]`
  - **A**: run `go run main.go migrate` to do migration
- **Q**: got error `Command 'air' not found`
  - **A**: install [air](//github.com/cosmtrek/air)
- **Q**: got error `Command 'replacer' not found`
  - **A**: install [replacer](//github.com/kokizzu/replacer)
- **Q**: got error `Command 'gomodifytags' not found`
  - **A**: install [goimports](//github.com/fatih/gomodifytags)
- **Q**: got error `Command 'farify' not found`
  - **A**: install [farify](//github.com/akbarfa49/farify)
- **Q**: where to put secret that I don't want to commit?
  - **A**: on `.env.override` file
- **Q**: got error `.env.override` no such file or directory
  - **A**: create `.env.override` file
- **Q**: got error `failed to stat the template: index.html`
  - **A**: run `cd svelte; npm run watch` at least once
- **Q**: got error `TarantoolConf) Connect: dial tcp 127.0.0.1:3301: connect: connection refused"`
  - **A**: run `docker-compose up`
- **Q**: got error `ClickhouseConf) Connect: dial tcp 127.0.0.1:9000: connect: connection refused`
  - **A**: run `docker-compose up`
- **Q**: got error `docker.errors.DockerException: Error while fetching server API version: ('Connection aborted.', FileNotFoundError(2, 'No such file or directory'))`
  - **A**: make sure docker service is up and running
- **Q**: what's normal flow of development?
  - **A**: 
      1. create new/modify model on `model/m[schema]/[schema]_tables.go` folder, create benchmark function to generate and migrate the tables in `RunMigration` function.
      2. run `./gen-orm.sh`, create helper function on `model/w[schema]/[rq|wc|sa][schema]/[schema]_helper.go` or 3rd party wrapper in `model/x[repo]/x[provider].go`
      3. create a role in `domain/[role].go` containing all business logic for that role
      4. write test in `domain/[role]_test.go` to make sure all business requirement are met
      5. generate domain routes `cd presentation; go get -bench=BenchmarkGenerateViews`, start web service `air web`
      6. write frontend on `svelte/`, start frontend service `cd svelte; npm run watch`
      7. generate frontend helpers `cd presentation; go get -bench=BenchmarkGenerateViews`
      8. write SSR if needed on `presentation/web_static.go`
- **Q**: how to add additional tech stack?
  - **A**: put on `docker-compose.yml` and add to `domain/0_main_test.go` so it would run on integration test. Create the `conf/[provider].go` and `model/x[repo]/x[provider].go` to wrap the 3rd party connector. 
- **Q**: want to change the generated views?
  - **A**: 
      1. change the template on `presentation/1_codegen_test.go`
      2. run `cd presentation; go get -bench=BenchmarkGenerateViews`
- **Q**: want to change generated ORM or schema migration has bug?
  - **A**: create a pull request to [gotro](//github.com/kokizzu/gotro) 
- **Q**: generated html have bug?
  - **A**: create a pull request to [svelte-mpa](//github.com/kokizzu/svelte-mpa)
- **Q**: where is the devlog?
  - **A**: on [youtube livestream](//www.youtube.com/@kokizzu/streams)
- **Q**: why secrets not encrypted?
  - **A**: it's ok for PoC phase, since it's listen to localhost
