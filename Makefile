install-dependencies:
	 go install github.com/jstemmer/go-junit-report@latest
	 go install gotest.tools/gotestsum@latest

run-test:
	gotestsum --format testname

run:
	go run ./cmd

docker-build:
	docker build -t property-finder-case-study-server-1 .

docker-run:
	docker run --name pproperty-finder-case-study-server-1 --env-file ./.env -p 8000:8000 -d property-finder-case-study-server-1

docker-develop:
	docker run --name property-finder-case-study-server-1 --env-file ./.env -p 8000:8000 -v `pwd`:`pwd` -w `pwd` -i -t -d property-finder-case-study-server-1

make-mocks:
	mockgen -source=internal/domain/user/repository_user/repository_user.go -destination=test_data/mocks/user_repository_mock.go -package=mocks
	mockgen -source=internal/domain/product/repository_product/repository_product.go -destination=test_data/mocks/product_repository_mock.go -package=mocks
	mockgen -source=internal/domain/order/repository_order/repository_order.go -destination=test_data/mocks/order_repository_mock.go -package=mocks
	mockgen -source=internal/domain/cart/repository_cart/repository_cart.go -destination=test_data/mocks/cart_repository_mock.go -package=mocks

make run:
	docker compose up -d