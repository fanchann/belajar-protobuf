test:
	go test -v -run=. github.com/fanchann/belajar-protobuf/tests

protoc-gen:
	cd proto && sudo protoc --proto_path=${PWD}/proto --go_out=. *.proto