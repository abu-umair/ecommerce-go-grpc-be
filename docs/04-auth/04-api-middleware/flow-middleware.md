
---

# 📂 2. Middleware Docs (`docs/04-auth/04-api-middleware`)

### `flow-middleware.md`
```markdown
# Flow Middleware

## Alur gRPC
1. Request masuk ke server.  
2. `ErrorMiddleware` → catch panic/error global.  
3. `AuthMiddleware` → cek token, cek blacklist, parse claims, set ke context.  
4. Handler → Service → Repository.  

## Laravel Analogi
- `ErrorMiddleware` → mirip `app/Exceptions/Handler.php`.  
- `AuthMiddleware` → mirip `Authenticate::class` atau JWT Middleware di Laravel.  
