# create:
# 	gk n s add
# 	sed -i "" 's/Foo(ctx context.Context, s string) (res string, err error)/Sum(ctx context.Context, a, b int64) (res int64, err error)/g' internal/app/add/service/service.go
# 	gk init add
# 	sed -i "" 's/return res, err/return a + b, err/g' internal/app/add/service/service.go
# 	gk add grpc add
# 	cd pb/add && ./compile.sh
# 	gk init grpc add
# 	gk new cmd add

add:
	curl -X POST localhost:8180/sum -d '{"a": 1, "b":1}'

gadd:
	grpcurl -plaintext -proto ./pb/add/add.proto -d '{"a": 1, "b":1}' localhost:8181 pb.Add.Sum

tic:
	curl -X POST localhost:9190/tic

tac:
	curl localhost:9190/tac	

# build:
# 	pack build add --env "GOOGLE_BUILDABLE=cmd/add/main.go" --builder gcr.io/buildpacks/builder:v1

# istio
a:
	curl -X POST localhost:80/api/v1/add/sum -d '{"a": 1, "b":1}'

b:
	grpcurl -plaintext -proto ./pb/add/add.proto -d '{"a": 1, "b":1}' localhost:443 pb.Add.Sum

c:
	curl -X POST localhost:80/api/v1/count/tic

d:
	curl localhost:80/api/v1/count/tac

x:
	helm install --name postgresql \
    --set global.postgresql.postgresqlDatabase=count \
    --set global.postgresql.postgresqlPassword=password \
    stable/postgresql