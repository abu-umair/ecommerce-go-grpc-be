### `03-jwt.go.md`
```markdown
# File: internal/entity/jwt/jwt.go
Fokus untuk middleware (parse claims).

```go
func (jc *JwtClaims) SetToContext(ctx context.Context) context.Context {
    return context.WithValue(ctx, JwtEntityContextKeyValue, jc)
}
```

## Laravel Analogi

SetToContext(ctx) → mirip inject user ke request()->attributes->set('user', $claims).