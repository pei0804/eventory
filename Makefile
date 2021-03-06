#! /usr/bin/make
#
# Targets:
# - clean     delete all generated files
# - generate  (re)generate all goagen-generated files.
# - build     compile executable
# - ae-build  build appengine
# - ae-dev    deploy to local (dev) appengine
# - ae-deploy deploy to appengine
#
# Meta targets:
# - all is the default target, it runs all the targets in the order above.
#

build: install appengine

install:
	@which direnv || go get -v github.com/zimbatm/direnv
	@direnv allow
	@which glide || go get -v github.com/Masterminds/glide
	@glide install

REPO:=github.com/tikasan/eventory

appengine:
	@which gorep || go get -v github.com/novalagung/gorep
	@gorep -path="./vendor/github.com/goadesign/goa" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./models" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./app" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./client" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./tool" \
          -from="context" \
          -to="golang.org/x/net/context"
	@gorep -path="./" \
          -from="../app" \
          -to="$(REPO)/app"
	@gorep -path="./" \
          -from="../client" \
          -to="$(REPO)/client"
	@gorep -path="./" \
          -from="../tool/cli" \
          -to="$(REPO)/tool/cli"

test:
	go test -v $(shell glide novendor)

##### goa ######

all: clean generate appengine

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger

generate:
	@goagen app     -d github.com/tikasan/eventory/design
	@goagen swagger -d github.com/tikasan/eventory/design
	@goagen client  -d github.com/tikasan/eventory/design

model:
	@rm -rf models
	@goagen --design=github.com/tikasan/eventory/design gen --pkg-path=github.com/goadesign/gorma


##### Database ######

DBNAME:=eventory
ENV:=development

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

##### Docker ######

docker/build: Dockerfile docker-compose.yml
	docker-compose build

docker/start:
	docker-compose up -d

docker/stop:
	docker-compose down

docker/logs:
	docker-compose logs

docker/clean:
	docker-compose rm

##### App engine ######

PROJECT:=eventory-test

deploy:
	goapp deploy -application $(PROJECT) ./server

rollback:
	appcfg.py rollback ./server -A $(PROJECT)

local:
	goapp serve ./server

school:
	goapp serve -port 8010 ./server
