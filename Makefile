run:
	go run main.go

protoc-gen:
	cd proto && sudo protoc --proto_path=${PWD}/proto --go_out=. *.proto
	