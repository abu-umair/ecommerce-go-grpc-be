
---

### `04-jwt.go.md`
```markdown
# File: internal/entity/jwt/jwt.go
Fokus: validasi JWT & ambil claims.

```go
// Claims JWT custom
type JwtClaims struct {
    jwt.RegisteredClaims
    Email    string `json:"email"`
    FullName string `json:"full_name"`
    Role     string `json:"role"`
}

// Parse token string jadi entity claims
func GetClaimsFromToken(token string) (*JwtClaims, error) {
    tokenClaims, err := jwt.ParseWithClaims(token, &JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
        if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil || !tokenClaims.Valid {
        return nil, utils.UnauthenticatedResponse()
    }

    if claims, ok := tokenClaims.Claims.(*JwtClaims); ok {
        return claims, nil
    }
    return nil, utils.UnauthenticatedResponse()
}
```

### Laravel Analogi

JwtClaims → mirip payload JWT (sub, email, role).

GetClaimsFromToken(token) → mirip JWTAuth::setToken($token)->getPayload().

Validasi algoritma HMAC → mirip JWT library Laravel yang cek signature.
