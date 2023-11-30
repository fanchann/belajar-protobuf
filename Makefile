run:
	go run main.go

protoc-gen:
	cd proto && sudo protoc --go_out=. *.proto
	