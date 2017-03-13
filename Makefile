all:
	cd cmd/hypnos && go build

install:
	cd cmd/hypnos && go install

macos-release:
	cd cmd/hypnos && go build
	cd cmd/hypnos && tar -zcvf ../../hypnos-macos.tar.gz hypnos && mv hypnos ../..

linux-release:
	docker build -f Dockerfile -t hypnos .
	docker run --name hypnos hypnos
	docker cp hypnos:/hypnos hypnos
	chmod +x hypnos
	tar -zcvf hypnos-linux-amd64.tar.gz hypnos
	docker rm hypnos

clean:
	-rm hypnos
	-rm cmd/hypnos/hypnos 
	-rm hypnos-macos.tar.gz
	-rm hypnos-linux-amd64.tar.gz
