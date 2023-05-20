GO = go
SRC_DIR = ./src
EXEC = monkey

.PHONY: all build run fmt test clean

all: build

build:
	$(GO) build -o $(EXEC) $(SRC_DIR)

run:
	$(GO) run $(SRC_DIR)

fmt:
	find $(SRC_DIR) -name "*.go" -exec gofmt -w {} \;

test:
		$(GO) test $(SRC_DIR)/...

clean:
	rm -f $(EXEC)
