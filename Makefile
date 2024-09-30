yaml_file = ./files/yml/library.local.yml
cmd_folder = ./cmd/library_app/
gorun = @go run

run:
	${gorun} ${cmd_folder}http -config ${yaml_file}

migration:
	${gorun} ${cmd_folder}migration -config ${yaml_file}

seed:
	${gorun} ${cmd_folder}seed -config ${yaml_file}

drop:
	${gorun} ${cmd_folder}drop -config ${yaml_file}

refresh: drop migration seed

scheduler:
	${gorun} ${cmd_folder}scheduler -config ${yaml_file}