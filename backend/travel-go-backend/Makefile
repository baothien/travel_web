.PHONY: run
run:
	cd ./services/$(s)/cmd/app && go run main.go
doc:
	cd ./services/$(s) && swag init -g /cmd/app/main.go --parseDependency --parseInternal --parseDepth 1
