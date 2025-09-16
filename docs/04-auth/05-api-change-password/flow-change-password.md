
---

### 📄 05-flow-change-password.md

```markdown
# Flow Change Password

## Alur gRPC
1. Client kirim request `ChangePasswordRequest { old_password, new_password, new_password_confirmation }`.  
2. Handler (`auth_handler.go`) → validasi request.  
3. Service (`auth_service.go`) →  
   - Cek konfirmasi password.  
   - Ambil user dari JWT token.  
   - Validasi password lama.  
   - Hash password baru.  
   - Update DB lewat repository.  
4. Repository (`auth_repository.go`) → jalankan query update.  
5. Response → return `ChangePasswordResponse { message: "Change password success" }`.  

---

## Analogi Laravel Flow
1. Client kirim request ke endpoint `/change-password`.  
2. `ChangePasswordRequest.php` (FormRequest) validasi input.  
3. `AuthController@changePassword` → terima request.  
4. `AuthService::changePassword` → proses logika utama.  
5. `User::update([...])` → update password di database.  
6. Return response JSON → `{ "message": "Change password success" }`.  
```