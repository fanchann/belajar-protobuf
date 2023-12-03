# Paths
PROTO_PATH = $(PWD)/proto
OUTPUT_DIR = $(PWD)

# Proto files

APP_SERVICES_PROTO = service/app_service.proto
JOBS_SERVICES_PROTO = jobs/job_service.proto

protoc-gen:
	protoc --proto_path=$(PROTO_PATH) --go_out=$(OUTPUT_DIR) $(PROTO_PATH)/*.proto
	protoc --proto_path=$(PROTO_PATH) --go_out=$(OUTPUT_DIR) $(PROTO_PATH)/$(APP_SERVICES_PROTO)
	protoc --proto_path=$(PROTO_PATH) --go_out=$(OUTPUT_DIR) $(PROTO_PATH)/$(JOBS_SERVICES_PROTO)
