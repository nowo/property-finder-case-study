test:
	go test ./...

run:
	go run ./cmd

docker-build:
	docker build -t property-finder-case-study-server-1 .

docker-run:
	docker run --name pproperty-finder-case-study-server-1 --env-file ./.env -p 8000:8000 -d property-finder-case-study-server-1

docker-develop:
	docker run --name property-finder-case-study-server-1 --env-file ./.env -p 8000:8000 -v `pwd`:`pwd` -w `pwd` -i -t -d property-finder-case-study-server-1

docker-test:
	docker exec -it property-finder-case-study-server-1 go test ./test/... -v


make-mocks:
	mockgen -source=internal/domain/user/repository_user/irepository_user.go -destination=test/mocks/user_repository_mock.go -package=mocks
	mockgen -source=internal/domain/product/repository_product/irepository_product.go -destination=test/mocks/product_repository_mock.go -package=mocks
	mockgen -source=internal/domain/order/repository_order/irepository_order.go -destination=test/mocks/order_repository_mock.go -package=mocks
	mockgen -source=internal/domain/cart/repository_cart/irepository_cart.go -destination=test/mocks/cart_repository_mock.go -package=mocks