linux_build:
	go build -o build/bin/app cmd/main/app.go

linux_run: linux_build
	./build/bin/app

win_build:
	go build -o build/bin/app.exe cmd/main/app.go

win_run: win_build
	./build/bin/app.exe
	
linux_clean:
	rm -rf build

win_clean:
	rd /s /q build