
---

### 📄 03-auth_service.go.md

```markdown
# auth_service.go (Change Password)

Service berisi logika bisnis utama untuk ganti password.

```go
func (as *authService) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	// 1. Cek apakah password baru cocok dengan konfirmasi
	if request.NewPassword != request.NewPasswordConfirmation {
		return &auth.ChangePasswordResponse{
			Base: utils.BadRequestResponse("New password is not matched"),
		}, nil
	}
```
Validasi new_password harus sama dengan new_password_confirmation.
```go
	// 2. Ambil JWT token dari context
	jwtToken, err := jwtentity.ParseTokenFromContext(ctx)
```

Ambil JWT dari metadata request.
```go
	claims, err := jwtentity.GetClaimsFromToken(jwtToken)
```

Ambil data user (id, email) dari token.
```go
	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
```

Cari user di database berdasarkan email dari token.
```go
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.OldPassword))
```

Validasi apakah old_password cocok dengan yang ada di DB.
```go
	hashedNewPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 10)
```

Hash password baru.
```go
	err = as.authRepository.UpdateUserPassword(ctx, user.Id, string(hashedNewPassword), user.FullName)
```

Simpan password baru ke DB lewat repository.
```go
	return &auth.ChangePasswordResponse{
		Base: utils.SuccessResponse("Change password success"),
	}, nil
```

Return response sukses.

# analogi
AuthController@changePassword → ambil user dari Auth::user().

Hash::check($oldPassword, $user->password) → cek password lama.

User::update(['password' => Hash::make($newPassword)]) → simpan password baru.

return response()->json(['message' => 'Change password success']).