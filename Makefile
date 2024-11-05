include /Users/randy.steven/others/library-go/Library-GO/files/env/.env
export

check_env:
	@echo $(ENV)

prod_env = "prod"
local_env = "local"
stg_env = "stg"
dev_env = "dev"

yaml_file = ./files/yml/library.staging.yml
cmd_folder = ./cmd/library_app/
gorun = @go run


ifeq ($(ENV), $(prod_env))
	yaml_file = $(PROD_YML)
else ifeq ($(ENV), $(stg_env))
	yaml_file = $(STG_YML)
else ifeq ($(ENV), $(dev_env))
	yaml_file = $(DEV_YML)
else ifeq($(ENV), $(local_env))
	yaml_file = $(LOCAL_YML)
else
	$(error unknown env)
endif

run:
	${gorun} ${cmd_folder}http -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

seed:
	${gorun} ${cmd_folder}seed -config ${yaml_file}

drop:
	${gorun} ${cmd_folder}drop -config ${yaml_file}

alter:
	${gorun} ${cmd_folder}alter -config ${yaml_file}

build:
	@go build -o /bin/http

refresh: drop migration seed

scheduler:
	${gorun} ${cmd_folder}scheduler -config ${yaml_file}

test:
	go test -coverprofile=coverage.out ./... ;    go tool cover -html=coverage.out