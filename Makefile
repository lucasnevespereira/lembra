.PHONY: run build clean

APP_NAME=lembra

create:
	bin/$(APP_NAME) create --title='$(title)' --message='$(message)' --sound='$(sound)' --time='$(time)'

list:
	bin/$(APP_NAME) list

delete-all:
	bin/$(APP_NAME) delete --all=true

build:
	go build -o bin/$(APP_NAME)

clean:
	rm bin/$(APP_NAME)