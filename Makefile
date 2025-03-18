LOCAL_BIN:=$(CURDIR)/bin

generate-api:
	mkdir -p $(LOCAL_BIN)
	if test ! -f "$(LOCAL_BIN)/gowsdl"; then GOBIN=$(LOCAL_BIN) go install github.com/hooklift/gowsdl/cmd/gowsdl@latest; fi

	$(LOCAL_BIN)/gowsdl -p toledoapi -o toledoapi.go $(CURDIR)/toledoapi/toledoapi.wsdl
