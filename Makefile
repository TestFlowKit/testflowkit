RUN_CMD = go run ./cmd/etoolse/main.go run

lint:
	golangci-lint run

run:
	${RUN_CMD}

unit_test:
	 go test ./...

run_e2e:
	${RUN_CMD} -f e2e/frontend.yml -c e2e/cli.yml

run_e2e_server:
	go run e2e/server/main.go



