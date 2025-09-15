
---

### `06-response.go.md`
```markdown
# File: internal/utils/response.go
Response helper.

```go
func SuccessResponse(message string) *common.BaseResponse {
    return &common.BaseResponse{
        StatusCode: 200,
        Message:    message,
    }
}
```

### Analogi Laravel
SuccessResponse("Logout success") → mirip return response()->json(['status' => 200, 'message' => 'Logout success']);.