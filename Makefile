.PHONY: run build clean

APP_NAME=lembra

run:
	go run *.go --title='$(title)' --message='$(message)' --sound='$(sound)'

#run:
#	go run *.go --title="Hello there" --message="this is some content" --sound="Glass"

build:
	go build -o $(APP_NAME)

clean:
	rm $(APP_NAME)