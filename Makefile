exp:
	export DBURL="postgres://postgres:feruza1727@localhost:5432/farmish?sslmode=disable"

mig-up:
	migrate -path migrations -database ${DBURL} -verbose up

mig-down:
	migrate -path migrations -database ${DBURL} -verbose down


mig-create:
	migrate create -ext sql -dir migrations -seq create_table

mig-insert:
	migrate create -ext sql -dir migrations -seq insert_table

swag_init:
	swag init -g api/handler.go -o api/docs