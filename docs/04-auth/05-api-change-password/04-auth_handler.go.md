
---

### 📄 04-auth_handler.go.md

```markdown
# auth_handler.go (Change Password)

Handler menghubungkan request gRPC → service.

```go
func (sh *authHandler) ChangePassword(ctx context.Context, request *auth.ChangePasswordRequest) (*auth.ChangePasswordResponse, error) {
	validationErrors, err := utils.CheckValidation(request)
```

Jalankan validasi request.

```go
	if validationErrors != nil {
		return &auth.ChangePasswordResponse{
			Base: utils.ValidationErrorResponse(validationErrors),
		}, nil
	}
```
Jika validasi gagal → return error response.



```go
res, err := sh.authService.ChangePassword(ctx, request)
```
Jika validasi lolos → teruskan ke service.

```go
	return res, nil

```
Return response dari service.


# analogi
Mirip AuthController@changePassword → menerima request lalu memanggil AuthService::changePassword().

Validasi mirip ChangePasswordRequest.php.