# auth.proto (Login)

```proto
syntax = "proto3";

package auth;

import "common/base_response.proto";
import "buf/validate/validate.proto";

option go_package = "github.com/abu-umair/ecommerce-go-grpc-be/pb/auth";

//? sama kayak kamu bikin AuthController dengan method register() dan login() di Laravel
service AuthService {
  rpc Register(RegisterRequest) returns (RegisterResponse);
  rpc Login(LoginRequest) returns (LoginResponse); //? menambahkan service login
}

//? LoginRequest = LoginRequest.php (FormRequest di Laravel)
message LoginRequest { 
  string email = 1 [(buf.validate.field).string = { email: true, min_len: 1, max_len: 255 }]; //? required|email|min:1|max:255
  string password = 2 [(buf.validate.field).string = { min_len: 1, max_len: 255 }];           //? required|string|min:1|max:255
}

//? LoginResponse = JsonResponse di Laravel
message LoginResponse { 
  common.BaseResponse base = 1; //? status, message
  string access_token = 2;      //? sama kayak token di Laravel JWT / Sanctum
}
```


