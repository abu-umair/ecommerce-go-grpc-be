
---

## 📖 `06-entity_user.go.md`

```markdown
# entity/user.go = Model (Laravel: User.php)
Struktur data User.

```go
package entity

import "time"

const (
	UserRoleCustomer = "customer" // Laravel: const ROLE_CUSTOMER = 'customer';
	UserRoleAdmin    = "admin"    // Laravel: const ROLE_ADMIN = 'admin';
)

type User struct {
	Id        string
	FullName  string
	Email     string
	Password  string
	RoleCode  string
	CreatedAt time.Time
	CreatedBy *string
	UpdatedAt time.Time
	UpdatedBy *string
	DeletedAt time.Time
	DeletedBy *string
	IsDeleted bool
}
```
---

## analogi Laravel
### Models\User.php
```php
// App\Models\User.php
class User extends Model {
    protected $fillable = [
        'id','full_name','email','password','role_code',
        'created_at','created_by','updated_at','updated_by',
        'deleted_at','deleted_by','is_deleted'
    ];
}

```