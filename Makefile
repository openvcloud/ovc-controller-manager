binary: build
	go build -o build/ovc-controller-manager

build:
	mkdir -p build

docker: binary Dockerfile
	cp Dockerfile build
	docker build -t openvcloud/ovc-controller-manager:master ./build
