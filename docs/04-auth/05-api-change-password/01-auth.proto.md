# auth.proto (Change Password)

File ini mendefinisikan request & response untuk fitur Change Password di gRPC.

```proto
message ChangePasswordRequest { 
  string old_password = 1 [(buf.validate.field).string = { min_len: 1, max_len: 255 }];
  string new_password = 2 [(buf.validate.field).string = { min_len: 1, max_len: 255 }];
  string new_password_confirmation = 3 [(buf.validate.field).string = { min_len: 1, max_len: 255 }];
}
```
old_password → Password lama user (untuk verifikasi).

new_password → Password baru.

new_password_confirmation → Konfirmasi password baru.
```proto

message ChangePasswordResponse { 
  common.BaseResponse base = 1;
}

base → response standar (mirip JsonResponse di Laravel).
```

Analogi Laravel
ChangePasswordRequest = ChangePasswordRequest.php (FormRequest untuk validasi).