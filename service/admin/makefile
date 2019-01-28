build:
	protoc -I.  --go_out=plugins=micro:$(GOPATH)/src/gin_admin/service/admin user/user.proto
	protoc -I.  --go_out=plugins=micro:$(GOPATH)/src/gin_admin/service/admin role/*.proto
