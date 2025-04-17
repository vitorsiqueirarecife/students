.PHONY: dev clean

REFLEX = reflex
MAIN_FILE = ./cmd/server/main.go
GO_FILES = $(shell find . -name "*.go")

dev:
	@echo "Iniciando o servidor Go com reflex..."
	$(REFLEX) -r '\.go$$' -- go run $(MAIN_FILE)
