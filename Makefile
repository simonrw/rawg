SOURCES := $(wildcard *.go)
BIN := rawg

all: $(BIN)


$(BIN): $(SOURCES)
	go build -o $@ $(SOURCES)

.PHONY: clean
clean:
	@rm -f $(BIN)

