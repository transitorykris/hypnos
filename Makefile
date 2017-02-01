all:
	go build

install:
	go install

macos-release:
	go build
	tar -zcvf hypnos-macos.tar.gz hypnos

linux-release:
	docker build -f Dockerfile -t hypnos .
	docker run --name hypnos hypnos
	docker cp hypnos:/hypnos hypnos
	chmod +x hypnos
	tar -zcvf hypnos-linux-amd64.tar.gz hypnos
	docker rm hypnos
