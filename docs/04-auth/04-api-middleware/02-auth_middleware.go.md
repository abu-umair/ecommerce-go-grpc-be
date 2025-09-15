
---

### `02-auth_middleware.go.md`
```markdown
# File: internal/grpcmiddleware/auth_middleware.go
Auth middleware untuk gRPC.

```go
func (am *authMiddleware) Middleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
    log.Println(info.FullMethod)

    // Jika method Login/Register → skip auth
    if info.FullMethod == "/auth.AuthService/Login" || info.FullMethod == "/auth.AuthService/Register" {
        return handler(ctx, req)
    }

    // Ambil token
    tokenStr, err := jwtentity.ParseTokenFromContext(ctx)
    if err != nil {
        return nil, err
    }

    // Cek blacklist
    _, ok := am.cacheService.Get(tokenStr)
    if ok {
        return nil, utils.UnauthenticatedResponse()
    }

    // Parse jwt jadi claims
    claims, err := jwtentity.GetClaimsFromToken(tokenStr)
    if err != nil {
        return nil, err
    }

    // Set claims ke context
    ctx = claims.SetToContext(ctx)

    // Lanjutkan ke handler
    res, err := handler(ctx, req)

    return res, err
}
```

## Laravel Analogi

info.FullMethod → mirip Route::currentRouteName().

ParseTokenFromContext(ctx) → mirip JWTAuth::parseToken()->authenticate().

cacheService.Get(token) → mirip Cache::has(token) (cek token blacklist).

claims.SetToContext(ctx) → mirip request()->merge(['user' => $claims]).