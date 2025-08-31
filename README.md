# Go gRpc Be - Ecommerce

## E-Commerce Project - Setup
### Inisiasi Projek Server gRPC Go
#### History Steps
1. Initial Project Go (isi seperti github)
```bash
go mod init github.com/abu-umair/ecommerce-go-grpc-be

```
2. Mendownload Library gRpc (otomatis download library)
```bash
go get google.golang.org/grpc

```

3. Run Server
```bash
go run main.go
```
4. membuat koleksi proto (didalam folder proto, kemudian membuat service proto)

5. Generate Proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative service/service.proto
```
6. Menggunakan .env, dan download package
```bash
go get github.com/joho/godotenv

```
7. Run Ulang Server
```bash
go run main.go
```
8. mengecek di Postman dengan pilihan gRpc

### Setup Database
#### History Steps
1. Setting supabase dan .env
2. menambahkan package driver postgres (https://github.com/lib/pq)
```bash
go get github.com/lib/pq

```
3. Run Server
```bash
go run main.go
```

### Setup Error Middleware
#### History Steps
1. Run Server
```bash
go run main.go
```

2. Run Server
```bash
go run main.go
```

### Setup Response Wrapper
#### History Steps
1. Generate base_response
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative common/base_response.proto
```

2. Generate ulang service proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative service/service.proto
```

3. Run Server
```bash
go run main.go
```

### Setup Validation Error
#### History Steps
1. menggunakan validate github (copas filenya ke proto\buf\validate\validate.proto)
```bash
https://github.com/bufbuild/protovalidate/blob/main/proto/protovalidate/buf/validate/validate.proto
```

2. Menambahkan validasi (dokumentasi dibawah ini)
```bash
https://buf.build/bufbuild/protovalidate/docs/main:buf.validate
```

3. Ketinggalan import di service

4. Generate ulang service proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative service/service.proto
```