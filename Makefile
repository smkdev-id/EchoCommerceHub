all: clean tidy test test-cover build-app build-run

tidy:
	@go mod tidy

run:
	@go run cmd/main.go

clean:
	@echo "make clean 🧽"
	@rm -rf /tmp/* 2> /dev/null

test:
	@go test tests/promotion_test.go -coverpkg=./internal/services/promotions -coverprofile=api/result_tests.cov && go tool cover -func api/result_tests.cov

test-cover:
	@go tool cover -html=api/result_tests.cov

build-app:
	@go build -o bin/promotion-app cmd/main.go

build-run:
	@ ./bin/promotion-app