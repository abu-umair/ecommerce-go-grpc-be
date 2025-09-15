# File: internal/grpcmiddleware/error_middleware.go
Global error handler.

```go
func ErrorMiddleware(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
    defer func() {
        // Tangkap panic
        if r := recover(); r != nil {
            log.Println(r)
            debug.PrintStack() // cetak stack trace
            err = status.Errorf(codes.Internal, "Internal Server Error")
        }
    }()

    res, err := handler(ctx, req) // jalankan handler asli

    if err != nil {
        log.Println(err)
        return nil, status.Error(codes.Internal, "Internal Server Error")
    }

    // Cek error status khusus
    if st, ok := status.FromError(err); ok {
        if st.Code() == codes.Unauthenticated {
            return nil, err
        }
    }

    return res, err
}
```

### Laravel Analogi

defer recover() → mirip try { } catch (\Exception $e) global di Laravel.

debug.PrintStack() → mirip Log::error($e).

status.Errorf(codes.Internal, ...) → mirip return response()->json(['message' => 'Internal Server Error'], 500).