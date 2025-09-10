package jwt

func GetClaimsFromToken(token string) 

//? kembalikan token tadi hingga menjadi entity jwt
tokenClaims, err := jwt.ParseWithClaims(jwtToken, &entity.JwtClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	if !tokenClaims.Valid {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}

	var claims *entity.JwtClaims
	if claims, ok = tokenClaims.Claims.(*entity.JwtClaims); !ok {
		return nil, status.Error(codes.Unauthenticated, "Unauthenticated")
	}