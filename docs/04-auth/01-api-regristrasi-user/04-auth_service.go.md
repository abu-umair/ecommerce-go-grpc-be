
---

## 📖 `04-auth_service.go.md`

```markdown
# service/auth_service.go = Service Layer (Laravel: Services/AuthService.php)
Tempat logika bisnis (business logic) dijalankan.

```go
package service

import (
	"context"
	"time"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/utils"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/auth"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error)
}

type authService struct {
	authRepository repository.IAuthRepository // Hubungkan ke repository (DB layer)
}

func (as *authService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	// Cek password confirmation
	if request.Password != request.PasswordConfirmation {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("Password is not matched"), // Laravel: return response()->json(["error" => "Password mismatch"])
		}, nil
	}

	// Cek user existing
	user, err := as.authRepository.GetUserByEmail(ctx, request.Email) // Laravel: User::whereEmail($email)->first()
	if err != nil {
		return nil, err
	}
	if user != nil {
		return &auth.RegisterResponse{
			Base: utils.BadRequestResponse("User already exist"),
		}, nil
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10) // Laravel: Hash::make()
	if err != nil {
		return nil, err
	}

	// Insert user baru
	newUser := entity.User{
		Id:        uuid.NewString(),
		FullName:  request.FullName,
		Email:     request.Email,
		Password:  string(hashedPassword),
		RoleCode:  entity.UserRoleCustomer,
		CreatedAt: time.Now(),
		CreatedBy: &request.FullName,
	}

	err = as.authRepository.InsertUser(ctx, &newUser) // Laravel: User::create()
	if err != nil {
		return nil, err
	}

	return &auth.RegisterResponse{
		Base: utils.SuccessResponse("User is registered"),
	}, nil
}

func NewAuthService(authRepository repository.IAuthRepository) IAuthService {
	return &authService{
		authRepository: authRepository,
	}
}
```
---

## analogi Laravel
### AuthService.php
```bash
// App\Services\AuthService.php
class AuthService {
    public function register(array $data) {
        if ($data['password'] !== $data['password_confirmation']) {
            throw new \Exception("Password not matched");
        }

        if (User::where('email', $data['email'])->exists()) {
            throw new \Exception("User already exist");
        }

        $data['password'] = Hash::make($data['password']);

        User::create($data);

        return "User is registered";
    }
}

```