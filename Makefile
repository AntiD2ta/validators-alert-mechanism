all: compile run compress

compile:
	go build -o ./bin/ .

build-linux:
	env GOOS=linux go build -o ./bin/ .

run:
	./bin/validators-alert-mechanism

compress:
	cp ./bin/validators-alert-mechanism .
	zip ./compressed/validators-alert-mechanism ./validators.json ./validators-alert-mechanism
	rm ./validators-alert-mechanism