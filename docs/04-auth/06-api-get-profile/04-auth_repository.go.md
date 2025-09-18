
---

### 📄 `04-auth_repository.go.md`
```markdown
# repository/auth_repository.go (GetProfile)

```go
func (ar *authRepository) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	//? Query DB berdasarkan email
	row := ar.db.QueryRowContext(ctx, "SELECT id, email, password, full_name, role_code, created_at FROM \"user\" WHERE email = $1 AND is_deleted = false", email)

	if row.Err() != nil {
		return nil, row.Err()
	}

	var user entity.User
	err := row.Scan(
		&user.Id,
		&user.Email,
		&user.Password,
		&user.FullName,
		&user.RoleCode,
		&user.CreatedAt,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil //? user tidak ditemukan
		}
		return nil, err
	}

	return &user, nil
}
```

## Analogi Laravel
```php
$user = User::where('email', $email)->where('is_deleted', false)->first();

```