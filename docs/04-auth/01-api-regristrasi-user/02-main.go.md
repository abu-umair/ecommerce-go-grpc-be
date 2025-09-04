# main.go = index.php + RouteServiceProvider laravel
Tempat semua service diregistrasi dan server gRPC dijalankan.

---

```go
package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/abu-umair/ecommerce-go-grpc-be/internal/handler"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/repository"
	"github.com/abu-umair/ecommerce-go-grpc-be/internal/service"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/auth"
	"github.com/abu-umair/ecommerce-go-grpc-be/pkg/database"
	"github.com/abu-umair/ecommerce-go-grpc-be/pkg/grpcmiddleware"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() { 
	ctx := context.Background()           // Context global, mirip $request di Laravel
	godotenv.Load()                      // Load file .env (sama seperti Laravel Dotenv)

	lis, err := net.Listen("tcp", ":50052") // Buka port 50052 untuk gRPC (mirip php artisan serve --port=8000)
	if err != nil { 
		log.Panicf("Error when listening %v", err) // Kalau gagal buka port → panic (error fatal)
	}

	db := database.ConnectDB(ctx, os.Getenv("DB_URI")) // Koneksi database
	log.Println("Database is connected")               // Log kalau DB sukses terhubung

	// Dependency Injection
	authRepository := repository.NewAuthRepository(db) // Repository = query builder / Eloquent
	authService := service.NewAuthService(authRepository) // Service = Business Logic
	authHandler := handler.NewAuthHandler(authService)    // Handler = Controller Laravel

	// Buat server gRPC
	serv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			grpcmiddleware.ErrorMiddleware, // Middleware global (mirip App\Http\Kernel Laravel)
		),
	)

	// Daftarkan service Auth ke server gRPC
	auth.RegisterAuthServiceServer(serv, authHandler) 
	// Laravel: Route::post('/auth/register', [AuthController::class, 'register']);

	// Aktifkan reflection (untuk debugging di dev environment saja)
	if os.Getenv("ENVIRONMENT") == "dev" {
		reflection.Register(serv) 
		log.Println("Reflection is registered") 
	}

	// Jalankan server
	log.Println("Server is running on :50052 port") // Info server aktif
	if err := serv.Serve(lis); err != nil {         // Mulai terima request gRPC
		log.Panicf("Server is error %v", err)       // Kalau gagal → panic
	}
}
```
---

## analogi Laravel
### index.php
```bash
$app = require_once __DIR__.'/../bootstrap/app.php';
$kernel = $app->make(Illuminate\Contracts\Http\Kernel::class);
$response = $kernel->handle(
    $request = Illuminate\Http\Request::capture()
);
$response->send();
$kernel->terminate($request, $response);

```

### routes/api.php
```bash
Route::post('/auth/register', [AuthController::class, 'register']);

```


