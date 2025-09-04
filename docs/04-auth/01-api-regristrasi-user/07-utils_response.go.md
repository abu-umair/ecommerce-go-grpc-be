
---

## 📖 `07-utils_response.go.md`

```markdown
# utils/response.go = Helper Response (Laravel: response()->json())
Standardisasi response.

```go
package utils

import "github.com/abu-umair/ecommerce-go-grpc-be/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message, // Laravel: return response()->json(["message" => $message], 200)
	}
}

func BadRequestResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 400,
		Message:    message,
		IsError:    true, // Laravel: return response()->json(["error" => $message], 400)
	}
}

func ValidationErrorResponse(validationErrors []*common.ValidationError) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode:      400,
		Message:         "Validation error",
		IsError:         true,
		ValidationError: validationErrors, // Laravel: return response()->json($validator->errors(), 400)
	}
}
```
---