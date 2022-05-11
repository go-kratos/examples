package v1

// 生成 proto grpc
//go:generate protoc --proto_path=. --proto_path=../../../third_party --proto_path=../../../ --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. ./*.proto

// 生成 proto http
//go:generate protoc --proto_path=. --proto_path=../../../third_party --proto_path=../../../ --go_out=paths=source_relative:. --go-http_out=paths=source_relative:. ./*.proto

// 生成 proto errors
//go:generate protoc --proto_path=. --proto_path=../../../third_party --proto_path=../../../ --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. ./*.proto

// 生成 swagger
//go:generate protoc --proto_path=. --proto_path=../../../third_party --proto_path=../../../ --openapiv2_out . --openapiv2_opt logtostderr=true --openapiv2_opt json_names_for_fields=true ./*.proto
