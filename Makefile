
build:
	mkdir -p bin/
	go build -o bin/terraform-provider-opal

testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

sweep:
	@echo "WARNING: This will destroy infrastructure. Use only in development accounts."
	go test ./opal -v -sweep=test $(SWEEPARGS) -timeout 2m
