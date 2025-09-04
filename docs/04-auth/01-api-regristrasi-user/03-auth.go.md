
---

## 📖 `03-auth.go.md`

```markdown
# handler/auth.go = AuthController
Menerima request dari client → validasi → kirim ke service → balikan response.

```go
package handler

import (
	"context"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/auth"
)

type authHandler struct {
	auth.UnimplementedAuthServiceServer // gRPC handler default

	authService service.IAuthService    // Dependency injection ke service
}

func (sh *authHandler) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	validationErrors, err := utils.CheckValidation(request) // Laravel: $request->validate()
	if err != nil {
		return nil, err
	}

	if validationErrors != nil {
		return &auth.RegisterResponse{
			Base: utils.ValidationErrorResponse(validationErrors), // Laravel: return response()->json($validator->errors())
		}, nil
	}

	// Proses register lewat service
	res, err := sh.authService.Register(ctx, request) // Laravel: AuthService::register()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewAuthHandler(authService service.IAuthService) *authHandler {
	return &authHandler{
		authService: authService,
	}
}
```

---

## analogi Laravel
### AuthController.php

```php
// App\Http\Controllers\AuthController.php
class AuthController extends Controller {
    public function register(Request $request) {
        $validated = $request->validate([
            'email' => 'required|email',
            'password' => 'required|min:6|confirmed',
        ]);

        return $this->authService->register($validated);
    }
}

```