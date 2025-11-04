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

5. Mendownload package berikut
```bash
go get github.com/bufbuild/protovalidate-go
```
jika gagal
```bash
go get buf.build/go/protovalidate
```

6. Run Server
```bash
go run main.go
```
7. Test Postman

8. Generate ulang base_response
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative common/base_response.proto
```

9. Run Server
```bash
go run main.go
```

9. Run ulang Server
```bash
go run main.go
```

## E-Commerce Project - Autentikasi
### API Registrasi User
#### History Steps
1. Generate auth.proto nya
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative auth/auth.proto

```

2. Setelah generate, lihat file di 
```bash
auth/auth_grpc.pb.go
auth/auth.pb.go
```
ke 2 file tersebut di save, kemudian dilihat ada erorr atau tidak, jika tidak ada maka aman


3. install bcrypt
```bash
go get golang.org/x/crypto/bcrypt
```

4. install UUID
```bash
go get github.com/google/uuid
```

5. jalankan tidy
```bash
go mod tidy
```

6. Run Server
```bash
go run main.go
```

### API Login
#### History Steps
1. Generate ulang auth.proto nya
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative auth/auth.proto

```

2. install golang jwt
```bash
go get github.com/golang-jwt/jwt/v5
```

3. Run Server
```bash
go run main.go
```

### API Logout
#### History Steps
1. Generate ulang auth.proto nya, karena ada perubahan pada auth.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative auth/auth.proto

```

3. install go-cache
```bash
go get github.com/patrickmn/go-cache
```

4. jalankan tidy
```bash
go mod tidy
```

5. Run Server
```bash
go run main.go
```

### Implementasi Middleware Autentikasi
#### History Step

1. Run Server
```bash
go run main.go
```

### API Change Password
#### History Step

1. Generate ulang auth.proto nya, karena ada perubahan pada auth.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative auth/auth.proto

```

2. Run Server
```bash
go run main.go
```

### API Get Profile
#### History Step

1. Copy file proto (timestamp)
```bash
https://github.com/protocolbuffers/protobuf/blob/main/src/google/protobuf/timestamp.proto
```

2. Generate ulang auth.proto nya, karena ada perubahan pada auth.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative auth/auth.proto
```

3. Install timestamp nya
```bash
go get google.golang.org/protobuf/types/known/timestamppb
```

4. Run Server
```bash
go run main.go
```

5. convert ke website untuk melihat hasil timestamp : 1756904649 (optional)
```bash
https://www.epochconverter.com/
```

### Integrasi API Login FE
#### History Steps
1. Run Server
```bash
go run main.go
```

2. juga Run gRPC web proxy
```bash
grpcwebproxy --backend_addr=localhost:50052 --server_bind_address=0.0.0.0 --server_http_debug_port=8080 --run_tls_server=false --backend_max_call_recv_msg_size=577659248 --allow_all_origins
```

## E-Commerce Project - Produk
### API Tambah Produk
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```
2. Run Server
```bash
go run main.go
```

## E-Commerce Project - Produk
### Implementasi Upload Gambar Produk
#### History Steps
1. Mendownload modul gofiber (dipilih karena mudah setupnya)
```bash
go get github.com/gofiber/fiber/v2
```

2. Run Server 
```bash
go run cmd/grpc/main.go
```

## E-Commerce Project - Produk
### API Detail Produk
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```

2. Run Server 
```bash
go run cmd/grpc/main.go
```

3. ambil ID dari table product, kemudian pastekan di postman (pada Detail Product)

## E-Commerce Project - Produk
### API Edit Produk
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```

2. Run Server 
```bash
go run cmd/grpc/main.go
```

3. Run Server Rest Api 
```bash
go run cmd/rest/main.go
```

## E-Commerce Project - Produk
### API Delete Produk
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```

2. Run Server 
```bash
go run cmd/grpc/main.go
```

3. Run Server Rest Api 
```bash
go run cmd/rest/main.go
```

## E-Commerce Project - Produk
### API List Produk
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```

2. Generate pagination.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative common/pagination.proto
```

## E-Commerce Project - Produk
### API List Produk Admin
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```

## E-Commerce Project - Produk
### API Highlight Produk
#### History Steps
1. Generate product.proto
```bash
protoc --go_out=./pb --go-grpc_out=./pb --proto_path=./proto --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative product/product.proto
```

## E-Commerce Project - Produk
### Setup Form Tambah Produk Admin FE
#### History Steps
1. Run Server
```bash
go run cmd/grpc/main.go
```

2. juga Run gRPC web proxy
```bash
grpcwebproxy --backend_addr=localhost:50052 --server_bind_address=0.0.0.0 --server_http_debug_port=8080 --run_tls_server=false --backend_max_call_recv_msg_size=577659248 --allow_all_origins
```

### Integrasi Tambah Produk Admin FE
#### History Steps

1. Run Server
```bash
go run cmd/grpc/main.go
```

2. juga Run gRPC web proxy
```bash
grpcwebproxy --backend_addr=localhost:50052 --server_bind_address=0.0.0.0 --server_http_debug_port=8080 --run_tls_server=false --backend_max_call_recv_msg_size=577659248 --allow_all_origins
``` 

3. Run Server Rest Api 
```bash
go run cmd/rest/main.go
```

### Integrasi List Produk Admin FE
#### History Steps