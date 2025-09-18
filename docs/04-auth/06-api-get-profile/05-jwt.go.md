
---

### 📄 `05-jwt.go.md`
```markdown
# jwt/jwt.go (GetProfile)

```go
// Mengambil claims dari context
func GetClaimsFromContext(ctx context.Context) (*JwtClaims, error) {
	claims, ok := ctx.Value(JwtEntityContextKeyValue).(*JwtClaims)
	if !ok {
		return nil, utils.UnauthenticatedResponse()
	}

	return claims, nil
}
```

### Penjelasan

ctx.Value(JwtEntityContextKeyValue) = mengambil data user dari context gRPC.

JwtClaims berisi: Subject (user_id), Email, FullName, Role.

## Analogi Laravel

Sama persis seperti:
```php
$user = Auth::user(); // data diambil dari token JWT
```