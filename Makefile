build:
	go build -o build/bin/app cmd/main/app.go

run: build
	./build/bin/app
	
clean:
	rm -rf build
