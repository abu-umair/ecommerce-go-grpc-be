
---

### `05-parse.go.md`
```markdown
# File: internal/entity/jwt/parse.go
Ambil token dari metadata gRPC.

```go
func ParseTokenFromContext(ctx context.Context) (string, error) {
    md, ok := metadata.FromIncomingContext(ctx)
    if !ok {
        return "", utils.UnauthenticatedResponse()
    }

    bearerToken, ok := md["authorization"]
    if !ok || len(bearerToken) == 0 {
        return "", utils.UnauthenticatedResponse()
    }

    // Format harus "Bearer <token>"
    tokenSplit := strings.Split(bearerToken[0], " ")
    if len(tokenSplit) != 2 || tokenSplit[0] != "Bearer" {
        return "", utils.UnauthenticatedResponse()
    }

    return tokenSplit[1], nil
}
```

### Laravel Analogi

metadata.FromIncomingContext → mirip request()->header('Authorization').

Validasi "Bearer <token>" → mirip Str::startsWith($header, 'Bearer').

Return token → mirip JWTAuth::setToken($header)->authenticate().