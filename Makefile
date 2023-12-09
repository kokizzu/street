
setup:
	go get -u -v github.com/kokizzu/gotro@latest
	go install github.com/cosmtrek/air@latest
	go install github.com/fatih/gomodifytags@latest
	go install github.com/kokizzu/replacer@latest
	go install github.com/akbarfa49/farify@latest
	go install golang.org/x/tools/cmd/goimports@latest
	#curl -fsSL https://get.pnpm.io/install.sh | bash -
	curl -fsSL https://bun.sh/install | bash

local-tarantool:
	docker exec -it street-tarantool1-1 tarantoolctl connect userT:passT@127.0.0.1:3301
	# box.space -- list all tables
	# box.execute [[ SELECT * FROM "users" LIMIT 1 ]]
	# \set language sql
	# \set delimiter ;

local-clickhouse:
	docker exec -it street-clickhouse1-1 clickhouse-client -u userC
	# SHOW TABLES -- list all tables
	# SELECT * FROM "actionLogs" LIMIT 1;

modtidy:
	sudo chmod -R a+rwx _tmpdb && go mod tidy

fixtags:
	# fix struct tags
	go generate domain

orm:
	# generate ORM
	./gen-orm.sh

views:
	# generate views and routes
	./gen-views.sh
