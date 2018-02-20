binary: build
	go build -o build/cloud-controller-manager

build:
	mkdir -p build

docker: binary Dockerfile
	cp Dockerfile build
	docker build -t openvcloud/cloud-controller-manager:master ./build
