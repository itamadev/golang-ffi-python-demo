build:
	$(shell cd ./ffi &&  go get) 
	$(shell cd ./ffi && go mod tidy)
	$(shell cd ./ffi && go build -buildmode=c-shared -o ../engine/lib.so)

run: build
	python main.py
