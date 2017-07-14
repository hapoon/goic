vendor-install:
	glide install
	
vendor-update:
	glide update

test:
	go test --cover
	go vet
	