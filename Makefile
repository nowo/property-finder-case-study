test:
	go test ./...

run:
	go run ./cmd

docker-build:
	docker build -t property-finder-go-homework .

docker-run:
	docker run --name property-finder-go-homework --env-file ./.env -p 8000:8000 -d property-finder-go-homework

docker-develop:
	docker run --name property-finder-go-homework --env-file ./.env -p 8000:8000 -v `pwd`:`pwd` -w `pwd` -i -t -d property-finder-go-homework

docker-test:
	docker exec -it property-finder-go-homework go test ./tests/... -v