.PHONY: verify
verify: 
	go mod tidy && go mod verify

.PHONY: test
.DEFAULT_GOAL := test
test: verify
	for example in `ls | grep example`; do echo "Running in $$example"; (cd $$example; go run .); done


