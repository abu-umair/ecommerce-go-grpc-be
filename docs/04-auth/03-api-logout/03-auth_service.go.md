
---

### `03-auth_service.go.md`
```markdown
# File: internal/service/auth_service.go
Fungsi Logout di Service Layer.

```go
func (as *authService) Logout(ctx context.Context, request *auth.LogoutRequest) (*auth.LogoutResponse, error) {
    // Ambil token dari metadata (header Authorization)
    jwtToken, err := jwtentity.ParseTokenFromContext(ctx) 
    if err != nil {
        return nil, err
    }

    // Parse token jadi claims (isi payload JWT)
    tokenClaims, err := jwtentity.GetClaimsFromToken(jwtToken) 
    if err != nil {
        return nil, err
    }

    // Masukkan token ke blacklist (cache memory)
    as.cacheService.Set(
        jwtToken, 
        "", 
        time.Duration(tokenClaims.ExpiresAt.Time.Unix()-time.Now().Unix())*time.Second,
    )

    // Return response sukses
    return &auth.LogoutResponse{
        Base: utils.SuccessResponse("Logout success"),
    }, nil
}
```
### Laravel Analogi
ParseTokenFromContext(ctx) → mirip JWTAuth::parseToken()->getPayload().

GetClaimsFromToken → mirip akses payload JWT ($payload['sub'], $payload['exp']).

cacheService.Set(token, "", expired) → mirip Cache::put(token, true, $expireAt).

Return LogoutResponse → mirip return response()->json(['message' => 'Logout success']);.