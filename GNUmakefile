default: testacc

# Run acceptance tests
.PHONY: testacc
testacc:
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 120m

init:
	mkdir -p ~/.terraform.d/plugins/sagewave.io/sagewave/jumpcloud/0.1.0/darwin_amd64/
install:
	go build -o terraform-provider-jumpcloud
	mv terraform-provider-jumpcloud ~/.terraform.d/plugins/sagewave.io/sagewave/jumpcloud/0.1.0/darwin_amd64