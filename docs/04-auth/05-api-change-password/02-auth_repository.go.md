
### 📄 02-auth_repository.go.md

```markdown
# auth_repository.go (Change Password)

Repository menangani query database untuk update password.

```go
func (ar authRepository) UpdateUserPassword(ctx context.Context, userId string, hashedPassword string, updatedBy string) error {
	_, err := ar.db.ExecContext(
		ctx,
		"UPDATE \"user\" SET password = $1, updated_at = $2, updated_by = $3 WHERE id = $4",
		hashedPassword,
		time.Now(),
		updatedBy,
		userId,
	)
}
```
ExecContext → jalankan SQL update ke database.

password = $1 → password baru yang sudah di-hash.

updated_at = $2 → timestamp update.

updated_by = $3 → siapa yang melakukan update.

WHERE id = $4 → hanya update user tertentu.

# Analogi
```php
User::where('id', $userId)->update([
  'password' => Hash::make($newPassword),
  'updated_at' => now(),
  'updated_by' => $updatedBy
]);
```
