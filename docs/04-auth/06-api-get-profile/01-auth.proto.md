# auth.proto (GetProfile)

```proto
//? sama kayak bikin AuthController dengan method getProfile(Request $request).
service AuthService {
  rpc GetProfile(GetProfileRequest) returns (GetProfileResponse);
}

// Request kosong (karena data diambil dari token JWT, bukan body request)
message GetProfileRequest { 

}

// Response berisi data user
message GetProfileResponse { 
  common.BaseResponse base = 1; //? mirip JsonResponse::success()
  string user_id = 2;           //? sama kayak $user->id
  string full_name = 3;         //? sama kayak $user->full_name
  string email = 4;             //? sama kayak $user->email
  string role_code = 5;         //? sama kayak $user->role
  google.protobuf.Timestamp member_since = 6; //? mirip created_at
}
```

## Analogi Laravel

GetProfileRequest = tidak butuh FormRequest (karena datanya dari JWT).

GetProfileResponse = return response()->json([...]) atau UserResource.

