
 # applet
 goctl api go --dir ./app/applet --api ./app/applet/applet.api --style go_zero

 # user
 goctl rpc protoc ./app/user/rpc/user.proto --go_out=./app/user/rpc --go-grpc_out ./app/user/rpc --zrpc_out ./app/user/rpc --style go_zero
goctl model mysql datasource --dir ./app/user/rpc/model --table user --cache true --url "root:lwj@1993@tcp(127.0.0.1:3306)/zhihu_go"