
---

### 📄 `02-handler-auth.go.md`
```markdown
# handler/auth_handler.go (GetProfile)

```go
// method handler GetProfile
func (sh *authHandler) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	//? langsung lempar ke service
	res, err := sh.authService.GetProfile(ctx, request)
	if err != nil {
		return nil, err
	}

	return res, nil
}
```
### Penjelasan

ctx = context gRPC, isinya metadata termasuk JWT token.

request = kosong (karena request tidak membawa body).

Handler hanya meneruskan ke authService.GetProfile.

## Analogi Laravel

Sama kayak AuthController@getProfile(Request $request) yang langsung memanggil AuthService::getProfile($request).