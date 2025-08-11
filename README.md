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

4. membuat koleksi proto (didalam folder proto, kemudian membuat service proto)

5. Generate Proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative service/service.proto

6. Menggunakan .env, dan download package
```bash
go get github.com/joho/godotenv

```
