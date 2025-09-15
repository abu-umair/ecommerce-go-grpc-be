
### `02-auth_handler.go.md`
```markdown
# File: internal/handler/auth.go
Fokus pada Logout.

```go
// Handler implementasi Logout
func (sh *authHandler) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
    // Validasi request (meskipun LogoutRequest kosong, sistem tetap cek)
    validationErrors, err := utils.CheckValidation(request)
    if err != nil {
        return nil, err
    }

    if validationErrors != nil {
        return &auth.LogoutResponse{
            Base: utils.ValidationErrorResponse(validationErrors),
        }, nil
    }

    // Proses Logout via service
    res, err := sh.authService.Logout(ctx, request)
    if err != nil {
        return nil, err
    }

    return res, nil
}

```

### Laravel Analogi
Logout(ctx, request) → mirip AuthController@logout(Request $request).

CheckValidation → mirip LogoutRequest di Laravel.

authService.Logout → mirip AuthService::logout($request).

Return LogoutResponse → mirip return response()->json(['message' => 'Logout success']).
