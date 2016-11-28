DBNAME:=eventory
ENV:=development

setup:
	which sql-migrate || go get github.com/rubenv/sql-migrate/...
    which scaneo || go get github.com/variadico/scaneo
    glide install

test:
    go test -v $(shell glide novendor)

build:
	glide install
	go build -o cmd/eventory/main cmd/eventory/main.go

migrate/init:
	mysql -u root -h localhost --protocol tcp -e "create database \`$(DBNAME)\`" -p

migrate/up:
	sql-migrate up -env=$(ENV)

migrate/down:
	sql-migrate down -env=$(ENV)

migrate/status:
	sql-migrate status -env=$(ENV)

migrate/dry:
	sql-migrate up -dryrun -env=$(ENV)

gen:
	cd model && go generate
