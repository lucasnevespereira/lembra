.PHONY: all run build clean create update delete list logs listen stop lint

APP_NAME=lembra

# CRUD
create:
	bin/$(APP_NAME) create --title='$(title)' --message='$(message)' --time='$(time)'

update:
	bin/$(APP_NAME) update --id='$(id)' --title='$(title)' --message='$(message)' --time='$(time)'

delete:
	bin/$(APP_NAME) delete --id='$(id)'

list:
	bin/$(APP_NAME) list

# DAEMON
logs:
	bin/$(APP_NAME) logs

listen:
	bin/$(APP_NAME) listen

stop:
	bin/$(APP_NAME) stop

# APP
build:
	go build -o bin/$(APP_NAME)

clean:
	rm bin/$(APP_NAME)

lint:
	golangci-lint run

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out | grep total | awk '{print "Total test coverage: " $$3}'
	rm coverage.out

test:
	go test -v ./...







