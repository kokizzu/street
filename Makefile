
local-tarantool:
	docker exec -it `docker ps | grep tarantool | cut -d ' ' -f 1` tarantoolctl connect userT:passT@localhost:3301
	# box.space -- list all tables
	# box.execute [[ SELECT * FROM "users" LIMIT 1 ]]
	# \set language sql
	# \set delimiter ;

local-clickhouse:
	docker exec -it `docker ps | grep clickhouse | cut -d ' ' -f 1` clickhouse-client --host localhost --port 9000 --user userC --password passC
	# SHOW TABLES -- list all tables
	# SELECT * FROM "actionLogs" LIMIT 1;