build:
	go build -o bin/task-tracker main.go

run: build
	./bin/task-tracker

clean:
	go clean
	rm -rf bin/