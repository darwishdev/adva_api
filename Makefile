LANG=en_US.UTF-8
SHELL=/bin/bash
.SHELLFLAGS=--norc --noprofile -e -u -o pipefail -c
# Include the main .env file
include config/state.env
# Construct the variable name based on STATE
CURRENT_STATE_FILE = config/$(STATE).env
# Include the appropriate .env file (e.g., dev.env or prod.env)
include $(CURRENT_STATE_FILE)

# Include the additional .env file
include config/shared.env
 

test:
	go test  -v -race  -cover ./... 

seed: 
	curl -X POST ${SEEDER_URL}/accounts/settings && curl -X POST ${SEEDER_URL}/accounts/permissions && curl -X POST ${SEEDER_URL}/accounts/roles && curl -X POST ${SEEDER_URL}/accounts/users && curl -X POST ${SEEDER_URL}/accounts/navigations 
 
seed_t:
	curl -X POST http://192.168.1.40:3000/users -d '{"test":true}' -H "Content-Type: application/json"
	
db_s:
	supabase start -x realtime 
mign : 
	supabase migration new $(name)
migu : 
	supabase migration up 
migd: 
	migrate -path common/db/migration -database ${DB_SOURCE} -verbose down 

cdb: 
	docker exec -it postgres  createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}
	
ddb :
	docker exec -it postgres  dropdb --username=${DB_USER}   ${DB_NAME}  --force
rdb:
	supabase db reset
rdbr:
	supabase db reset --linked

rdbs:
	make rdb seed
rdb_t:
	make  ddb cdb migu seed_t
run:
	go run main.go



mock:
	mockgen -package mockdb -destination common/db/mock/store.go github.com/darwishdev/bzns_pro_api/common/db/gen Store
sqlc :
	rm -rf common/db/gen/*.sql.go && sqlc generate


buf:
	rm -rf common/pb/* && buf generate



buf_push_g:
	cd common/proto && git add . && git commit -m "sync" && git push -u origin main
buf_push:
	cd common/proto && buf push



buf_u:
	buf mod update


dtag:
	docker tag bzns_pro exploremelon/bzns_pro:${v}

dpush:
	docker push  exploremelon/bzns_pro:${v}

run_bg:
	make run  > /dev/null 2>&1 & && disown


gen:
	rm -rf lib/pb/*  && protoc -I=proto/mln_api_protos --dart_out=grpc:lib/pb proto/mln_api_protos/bznspro/v1/*.proto proto/mln_api_protos/bznspro/v1/*/*.proto  proto/mln_api_protos/google/protobuf/*.proto
