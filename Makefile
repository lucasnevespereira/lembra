.PHONY: run build clean

APP_NAME=lembra

run:
	go run *.go --title='$(title)' --message='$(message)' --sound='$(sound)' --time='$(time)'

#run:
#	go run *.go --title="Hello there" --message="this is some content" --sound="Glass"
# ./bin/lembra create --title="Hello" --message="you have a meeting" --time="14h08"

build:
	go build -o bin/$(APP_NAME)

clean:
	rm bin/$(APP_NAME)