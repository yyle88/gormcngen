COVERAGE_DIR ?= .coverage

# cp from: https://github.com/yyle88/gormcnm/blob/6854ab1ff2bf1824add8eff3730c9fa4bd0b71fb/Makefile#L4
test:
	@-rm -r $(COVERAGE_DIR)
	@mkdir $(COVERAGE_DIR)
	make test-with-flags TEST_FLAGS='-v -race -covermode atomic -coverprofile $$(COVERAGE_DIR)/combined.txt -bench=. -benchmem -timeout 20m'

test-with-flags:
	@go test $(TEST_FLAGS) ./...
