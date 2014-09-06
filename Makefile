all: cli-weather


cli-weather:
	cd cli && go build $@.go 

clean:
	@if [ -e cli/cli-weather ];\
	then \
	echo "Removing cli/cli-weather"; \
	rm cli/cli-weather; \
	fi
