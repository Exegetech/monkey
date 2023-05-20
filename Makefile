GO = go
SRC_DIR = ./src
TEST_DIR = ./tests
EXEC = monkey

.PHONY: all build run fmt test clean

all: build

build:
	$(GO) build -o $(EXEC) $(SRC_DIR)

run:
	$(GO) run $(SRC_DIR)

fmt:
	find . -name "*.go" -exec gofmt -w {} \;

test:
		$(GO) test $(TEST_DIR)/...

clean:
	rm -f $(EXEC)
