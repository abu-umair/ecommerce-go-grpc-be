# Flow Registrasi User (dengan .proto)
## 1. Definisi kontrak di .proto

// Sebelum ada Controller atau Repository, kita tulis dulu service, request, dan response di .proto.
// Laravel tidak ada step ini, langsung bikin Route + Controller.

---

## 2. Generate kode Go

// Dari .proto, gRPC generate otomatis auth.pb.go dan auth_grpc.pb.go.
// Ini mirip artisan di Laravel yang generate Controller & Request class.

---

## 3. Handler Layer

// gRPC: implementasi AuthService.Register() di auth.go.
// Laravel: AuthController@register.

---

## 4. Validasi

// gRPC: dilakukan otomatis dari definisi (buf.validate.field) atau manual di utils.CheckValidation.
// Laravel: otomatis via RegisterRequest.

---

## 5. Service Layer

// gRPC: authService.Register() → cek email, hash password, dll.
// Laravel: AuthService::register().

---

## 6. Repository Layer

// gRPC: authRepository.InsertUser().
// Laravel: User::create([...]).

---

## 7. Response

// gRPC: return RegisterResponse dengan common.BaseResponse.
// Laravel: return response()->json([...]).

--- 

# Diagram Singkat

## gRPC
Client → .proto (kontrak) → Handler → Service → Repository → Database → Response

## Laravel
Client → Route → Controller → Service (opsional) → Model → Database → Response