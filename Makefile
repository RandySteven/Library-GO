yaml_file = ./files/yml/library.local.yml
cmd_folder = ./cmd/library_app/
gorun = @go run

run:
	${gorun} ${cmd_folder}http -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

seed:
	${gorun} ${cmd_folder}seed -config ${yaml_file}
