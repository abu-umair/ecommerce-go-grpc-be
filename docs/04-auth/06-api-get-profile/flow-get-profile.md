# Flow GetProfile API (gRPC)

## Alur Eksekusi
1. **Client** memanggil `GetProfile` ke `AuthService` di `auth.proto`.
2. gRPC memanggil **Handler** (`authHandler.GetProfile`).
3. Handler meneruskan request ke **Service** (`authService.GetProfile`).
4. Service:
   - Mengambil **JWT token** dari context.
   - Mengecek claims JWT di `jwt.go`.
   - Mengambil data user dari **Repository** (`GetUserByEmail`).
   - Menyusun response.
5. Repository menjalankan query ke DB untuk mencari user berdasarkan email.
6. Response dikembalikan step by step → Service → Handler → Client.

---

## Analogi Laravel
- `auth.proto` = **AuthController** dengan method `getProfile(Request $request)`.
- `handler` = **Controller** yang menerima request, validasi dasar, lalu lempar ke Service.
- `service` = **Service layer** / **Business logic** (kayak `AuthService` di Laravel).
- `repository` = **Eloquent/Model** (`User::where('email', $email)->first()`).
- `jwt.go` = **Middleware JWT** Laravel (`auth:api` guard).

Jadi alurnya mirip:
Route → Controller → Service → Repository → DB
↓
JWT Middleware