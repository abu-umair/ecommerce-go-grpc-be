
---

## 📖 `05-auth_repository.go.md`

```markdown
# repository/auth_repository.go = Model/Query Builder (Laravel: User::query())
Tempat query database.

```go
package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/entity"
)

type IAuthRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	InsertUser(ctx context.Context, user *entity.User) error
}

type authRepository struct {
	db *sql.DB
}

func (ar *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	row := ar.db.QueryRowContext(ctx, "SELECT id, email, password, full_name FROM \"user\" WHERE email = $1 AND is_deleted = false", email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user entity.User
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.FullName)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil // Laravel: User::whereEmail($email)->first() ?? null
		}
		return nil, err
	}

	return &user, nil
}

func (ar *authRepository) InsertUser(ctx context.Context, user *entity.User) error {
	_, err := ar.db.ExecContext(
		ctx,
		`INSERT INTO "user" (id, full_name, email, password, role_code, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted)
		 VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		user.Id,
		user.FullName,
		user.Email,
		user.Password,
		user.RoleCode,
		user.CreatedAt,
		user.CreatedBy,
		user.UpdatedAt,
		user.UpdatedBy,
		user.DeletedAt,
		user.DeletedBy,
		user.IsDeleted,
	)
	return err
}

func NewAuthRepository(db *sql.DB) IAuthRepository {
	return &authRepository{db: db}
}
```
---

## analogi Laravel
### Models\User.php
```php
// App\Models\User.php
class User extends Model {
    protected $fillable = ['id','full_name','email','password','role_code','created_at','created_by'];
}

```

### Example
```php
User::where('email', $email)->first();
User::create($data);


```