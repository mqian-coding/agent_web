Generate the proto files:

``protoc -I=proto --go_out=./proto/gen --go_opt=paths=source_relative proto/*.proto
``