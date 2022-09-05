CONFIG_PATH=${HOME}/.dist-service/

.PHONY: init
init:
	mkdir -p ${CONFIG_PATH}



.PHONY: gencert
gencert:
	cfssl gencert\
			-initca test/ca-csr.json | cfssljson -bare ca 
	cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config = test/ca-config.json \
			-profile=server \
			test/server-csr.json | cfssl -bare server 
	cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config = test/ca-config.json \
			-profile=client \
			-cn="root"\
			test/client-csr.json | cfssl -bare root-client
	cfssl gencert \
			-ca=ca.pem \
			-ca-key=ca-key.pem \
			-config = test/ca-config.json \
			-profile=client \
			-cn="nobody" \
			test/client-csr.json | cfssl -bare nobody-client

	mv *.pem *.csr ${CONFIG_PATH}


.PHONY: test
test:
	$(CONFIG_PATH)/policy.csv $(CONFIG_PATH)/model.conf 
	go test -race ./...

.PHONY: compile
compile:
	   protoc api/v1/*.proto \
	   --go_out=. \
	   --go-grpc_out=. \
	   --go-grpc_opt=paths=source_relative \
	   --go_opt=paths=source_relative \
	   --proto_path=.



$(CONFIG_PATH)/model.conf: cp model.conf $(CONFIG_PATH)/model.conf


$(CONFIG_PATH)/policy.csv: cp policy.csv $(CONFIG_PATH)/policy.csv

TAG ?=0.0.1

build-docker:
	docker build -t github.com/upalchowdhury/dist-service:$(TAG) .

