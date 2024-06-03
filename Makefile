exp:
	export DBURL="postgres://mrbek:QodirovCoder@localhost:5432/farmish?sslmode=disable"

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

	
gen-proto:
	protoc --go_out=services/ \
		--go-grpc_out=services/ \
		services/protos/*.proto

# mig-delete:
# 	rm -r db/migrations
# \