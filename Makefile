GO_CMD=go

tools:
	$(GO_CMD) install \
		github.com/fatih/errwrap/ \
		github.com/google/wire/cmd/wire \

db:
	docker run -p 27017:27017 -it --rm mongo
