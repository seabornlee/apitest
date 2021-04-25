.PHONY: test test-examples fmt vet

test:
	bash -c 'diff -u <(echo -n) <(go fmt $(go list ./...))'
	go vet ./...
	go test ./... -v -covermode=atomic -coverprofile=coverage.out

test-examples:
	cd examples && go test -v ./... && \
	cd sequence-diagrams-with-sqlite-database && make test && cd ..
