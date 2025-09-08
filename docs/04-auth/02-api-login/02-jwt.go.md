```go
package entity

import "github.com/golang-jwt/jwt/v5"

//? Sama kayak Laravel JWT payload (isi token)
type JwtClaims struct {
	jwt.RegisteredClaims
	Email    string `json:"email"`     //? user email
	FullName string `json:"full_name"` //? nama lengkap user
	Role     string `json:"role"`      //? role user (admin/customer)
}

//? Laravel analogy:
//? kalau pakai tymon/jwt-auth → isi token (sub, email, role, exp, iat)

```