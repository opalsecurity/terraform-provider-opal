
build:
	mkdir -p bin/
	go build -o bin/terraform-provider-opal

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m
