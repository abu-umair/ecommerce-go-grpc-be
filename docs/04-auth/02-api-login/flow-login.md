# Flow Login (gRPC vs Laravel)

Dokumen ini menjelaskan alur proses **Login User** pada aplikasi gRPC (Go) dan analoginya dengan Laravel.

---

## 1. auth.proto
- **Fungsi:** Mendefinisikan kontrak service `Login(LoginRequest) returns (LoginResponse)`.  
- **Kenapa penting:** Client & server harus sepakat struktur request/response.  
- **Laravel analogy:** Sama seperti bikin `routes/api.php` + `LoginRequest.php` (FormRequest untuk validasi).

---

## 2. Generate kode (auth.pb.go)
- **Fungsi:** gRPC generate interface `AuthServiceServer` dengan method `Login()`.  
- **Laravel analogy:** Artisan otomatis generate class Request atau Controller (`php artisan make:request LoginRequest`).

---

## 3. Handler Layer (auth_handler.go)
- **Fungsi:** Implementasi method `Login()` yang dipanggil client.  
- **Tugas:** meneruskan request ke `authService.Login()`.  
- **Laravel analogy:** Sama seperti `AuthController@login`.

---

## 4. Service Layer (auth_service.go)
- **Fungsi:** Berisi business logic login.  
  - Cari user by email → `authRepository.GetUserByEmail()`.  
  - Validasi password (bcrypt).  
  - Generate JWT token.  
- **Laravel analogy:** Sama seperti membuat `AuthService::login()` yang dipanggil dari Controller.

---

## 5. Entity (entity/jwt.go)
- **Fungsi:** Struktur `JwtClaims` untuk isi token (sub, email, role, expired, dll).  
- **Laravel analogy:** Sama dengan payload JWT yang otomatis dibentuk package `tymon/jwt-auth` atau `Laravel Sanctum`.

---

## 6. Repository Layer (auth_repository.go)
- **Fungsi:** Query database (`SELECT * FROM users WHERE email = ?`).  
- **Laravel analogy:** Sama dengan `User::where('email', $request->email)->first()`.

---

## 7. Response (LoginResponse)
- **Fungsi:** Return message + access_token.  
- **Laravel analogy:** Sama dengan:

```php
return response()->json([
    'message' => 'Login successful',
    'access_token' => $token,
]);
```