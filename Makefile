all: promotion-tidy promotion-tests promotion-tests-cover promotion-build-app promotion-build-run

promotion-tidy:
	@go mod tidy

promotion-run:
	@go run cmd/main.go

promotion-tests:
	@go test tests/promotion_test.go -coverpkg=./internal/app/services -coverprofile=api/result_tests.cov && go tool cover -func api/result_tests.cov

promotion-tests-cover:
	@go tool cover -html=api/result_tests.cov

promotion-build-app:
	@go build -o api/promotion-app cmd/main.go

promotion-build-run:
	@ ./api/promotion-app
