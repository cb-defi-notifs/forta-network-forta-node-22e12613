containers:
	docker build -t forta-network/forta-node -f Dockerfile.node .

containers-dev:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o forta-node cmd/node/main.go
	DOCKER_BUILDKIT=1 docker build --no-cache --network=host -t forta-network/forta-node -f Dockerfile.buildkit.dev.node .

main:
	docker build -t build-forta -f Dockerfile.cli .
	docker create --name build-forta build-forta
	docker cp build-forta:/main forta
	docker rm -f build-forta
	chmod 755 forta

.PHONY: mocks
mocks:
	mockgen -source clients/interfaces.go -destination clients/mocks/mock_clients.go
	mockgen -source clients/ratelimiter/rate_limiter.go -destination clients/ratelimiter/mocks/mock_rate_limiter.go
	mockgen -source services/components/registry/registry.go -destination services/components/registry/mocks/mock_registry.go
	mockgen -source store/registry.go -destination store/mocks/mock_registry.go
	mockgen -source services/storage/ipfs.go -destination services/storage/mocks/mock_ipfs.go
	mockgen -source store/scanner_release.go -destination store/mocks/mock_scanner_release.go
	mockgen -source services/components/metrics/lifecycle.go -destination services/components/metrics/mocks/mock_lifecycle.go
	mockgen -source services/components/lifecycle/bot_pool.go -destination services/components/lifecycle/mocks/mock_bot_pool.go
	mockgen -source services/components/lifecycle/bot_manager.go -destination services/components/lifecycle/mocks/mock_bot_manager.go
	mockgen -source services/components/lifecycle/bot_monitor.go -destination services/components/lifecycle/mocks/mock_bot_monitor.go
	mockgen -source services/components/containers/bot_client.go -destination services/components/containers/mocks/mock_bot_client.go
	mockgen -source services/components/containers/image_cleanup.go -destination services/components/containers/mocks/mock_image_cleanup.go
	mockgen -source services/components/botio/sender.go -destination services/components/botio/mocks/mock_sender.go
	mockgen -source services/components/botio/bot_client.go -destination services/components/botio/mocks/mock_bot_client.go
	mockgen -source services/components/botio/bot_client_factory.go -destination services/components/botio/mocks/mock_bot_client_factory.go
	mockgen -source clients/agentgrpc/dialer.go -destination clients/agentgrpc/mocks/mock_dialer.go
	mockgen -source clients/agentgrpc/client.go -destination clients/agentgrpc/mocks/mock_client.go
	mockgen -source services/jwt-provider/provider/jwt.go -destination services/jwt-provider/provider/mocks/mock_jwt.go


test:
	go test -v -count=1 ./... -coverprofile=coverage.out

.PHONY: coverage
coverage:
	go tool cover -func=coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'

coverage-func:
	go tool cover -func=coverage.out

coverage-html:
	go tool cover -html=coverage.out -o=coverage.html

perf-test:
	go test ./... -tags=perf_test

MOCKREG = $$(pwd)/tests/e2e/misccontracts/contract_mock_registry
MULTICALL = $$(pwd)/tests/e2e/misccontracts/contract_multicall

.PHONY: e2e-test-contracts
e2e-test-contracts:
	solc --bin --abi -o $(MOCKREG) --include-path . --base-path $(MOCKREG) --overwrite --input-file $(MOCKREG)/MockRegistry.sol
	abigen --out $(MOCKREG)/mock_registry.go --pkg contract_mock_registry --type MockRegistry --abi $(MOCKREG)/MockRegistry.abi --bin $(MOCKREG)/MockRegistry.bin
	solc --bin --abi -o $(MULTICALL) --include-path . --base-path $(MULTICALL) --overwrite --input-file $(MULTICALL)/Multicall.sol
	abigen --out $(MULTICALL)/multicall.go --pkg contract_multicall --type Multicall --abi $(MULTICALL)/Multicall.abi --bin $(MULTICALL)/Multicall.bin

.PHONY: e2e-test-deps
e2e-test-deps:
	./tests/e2e/deps-start.sh

.PHONY: e2e-test
e2e-test:
	./tests/e2e/build.sh

	cd tests/e2e && E2E_TEST=1 go test -v -count=1 .

run:
	go build -o forta . && ./forta --passphrase 123

build-local: ## Build for local installation from source
	./scripts/build-for-local.sh

build-remote: ## Try the "remote" containers option for build
	./scripts/build-for-release.sh disco-dev.forta.network

.PHONY: install
install: build-local ## Single install target for local installation
	cp forta /usr/local/bin/forta

.PHONY: update-core
update-core:
	go get github.com/forta-network/forta-core-go && go mod tidy
