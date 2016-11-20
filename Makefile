DBNAME:=eventory
ENV:=development

setup:
	go get github.com/rubenv/sql-migrate/...
	go get gopkg.in/yaml.v1
	go get github.com/go-sql-driver/mysql
	go get github.com/yterajima/go-dtf
	go get -v ./...

build:
	go build -o cmd/eventory/main cmd/eventory/main.go

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)