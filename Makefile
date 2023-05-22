.PHONY: run build clean

APP_NAME=lembra

run:
	go run *.go --title='$(title)' --message='$(message)' --sound='$(sound)' --time='$(time)'

build:
	go build -o bin/$(APP_NAME)

clean:
	rm bin/$(APP_NAME)