IMAGE_SYNC=eth-sync:latest
IMAGE_CHECK=eth-check:latest

dev-sync:
	docker build -f Dockerfile -t ${IMAGE_SYNC} sync.go

dev-check:
	docker build -f Dockerfile-check -t ${IMAGE_CHECK} check.go

run-sync:
	docker run --name eth-sync -d ${IMAGE_SYNC}

run-check:
	docker run --name eth-check -d ${IMAGE_CHECK}

sync:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o eth-sync sync.go

check:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o eth-check check.go