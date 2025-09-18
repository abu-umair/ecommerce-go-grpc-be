
---

### 📄 `03-auth_service.go.md`
```markdown
# service/auth_service.go (GetProfile)

```go
func (as *authService) GetProfile(ctx context.Context, request *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	//* Get data token dari context
	claims, err := jwtentity.GetClaimsFromContext(ctx)
	if err != nil {
		return nil, err
	}

	//* Ambil data user dari DB lewat repository
	user, err := as.authRepository.GetUserByEmail(ctx, claims.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return &auth.GetProfileResponse{
			Base: utils.BadRequestResponse("User doesn't exist"),
		}, nil
	}

	//* Susun response
	return &auth.GetProfileResponse{
		Base:        utils.SuccessResponse("Get Profile success"),
		UserId:      claims.Subject,
		FullName:    claims.FullName,
		Email:       claims.Email,
		RoleCode:    claims.Role,
		MemberSince: timestamppb.New(user.CreatedAt),
	}, nil
}
```
### Penjelasan

Ambil JWT claims (id, email, full_name, role).

Cari user di DB berdasarkan email.

Jika user ada → kembalikan response dengan data profil.

Jika user tidak ada → return error response.

## Analogi Laravel

jwtentity.GetClaimsFromContext(ctx) = Auth::user().

authRepository.GetUserByEmail = User::where('email', $email)->first().

Return response = return new UserResource($user).