all: compile run compress

compile:
	go build -o ./bin/ .

build-linux:
	env GOOS=linux go build -o ./bin/ .
	@make compress

run:
	./bin/validators-alert-mechanism

compress:
	@rm ./compressed/validators-alert-mechanism.zip
	@cp ./bin/validators-alert-mechanism .
	@zip ./compressed/validators-alert-mechanism ./validators-alert-mechanism
	@rm ./validators-alert-mechanism
