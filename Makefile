pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge

run:
	go run cmd/main.go

create_proto_submodule:
	git submodule add git@github.com:Asliddin3/Proto-Submodule-Product-servise.git

run_script:
	./script/gen-proto.sh

swag:
	swag init -g ./api/router.go -o api/docs
# user:
    # go run home/go/src/gitlab.com/grpc-first/user_service/cmd/main.go
	
# product:
# 	go run ~/go/src/gitlab.com/grpc-first/product-service/cmd/main.go 

# payment:
# 	go run ~/go/src/gitlab.com/grpc-first/payment-service/cmd/main.go 

# store:
# 	go run ~/go/src/gitlab.com/grpc-first/store-service/cmd/main.go 

	