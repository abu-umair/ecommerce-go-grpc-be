# File: proto/auth/auth.proto

```proto
// Service gRPC untuk Auth, mirip Controller di Laravel.
service AuthService {
  rpc Register(RegisterRequest) returns (RegisterResponse);
  rpc Login(LoginRequest) returns (LoginResponse);
  rpc Logout(LogoutRequest) returns (LogoutResponse); // Tambahan Logout
}

// Request Logout (kosong, karena cukup header Authorization saja).
message LogoutRequest { }

// Response Logout → return BaseResponse (status + message).
message LogoutResponse { 
  common.BaseResponse base = 1;
}
```
### Laravel Analogi

rpc Logout(...) → mirip bikin method logout() di AuthController.

LogoutRequest → mirip LogoutRequest.php (FormRequest kosong, hanya butuh token dari header).

LogoutResponse → mirip JsonResponse di Laravel (['status' => 200, 'message' => 'Logout success']).
