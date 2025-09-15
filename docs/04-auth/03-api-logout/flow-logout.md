# Flow Logout

## Alur gRPC
1. Client kirim request `LogoutRequest` → gRPC server.  
2. Handler (`auth.go`) validasi request.  
3. Service (`auth_service.go`) ambil token dari metadata via `ParseTokenFromContext`.  
4. Service parse token jadi claims via `jwt.go`.  
5. Token dimasukkan ke blacklist (cache memory).  
6. Kirim response sukses (`response.go`).  

## Laravel Analogi
- **Route**: `POST /logout`.  
- **Controller**: `AuthController@logout`.  
- **Ambil token**: `JWTAuth::parseToken()->getPayload()`.  
- **Blacklist**: simpan token di `Cache/Redis`.  
- **Response**: `return response()->json(['message' => 'Logout success']);`.
