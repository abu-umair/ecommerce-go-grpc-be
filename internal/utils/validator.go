package utils

import (
	"errors"

	"buf.build/go/protovalidate"
	"github.com/abu-umair/ecommerce-go-grpc-be/pb/common"
	"google.golang.org/protobuf/proto"
)

func CheckValidation(req proto.Message) ([]*common.ValidationError, error) {
	if err := protovalidate.Validate(req); //?validator.Validate(request) cek request sesuai aturan di .proto.
	err != nil {                           //?jika ada error maka akan mereturn error Analogi Laravel → mirip kalau kita bikin custom validator class lalu panggil Validator::make($data, $rules).
		var validationError *protovalidate.ValidationError //?Kalau gagal → return *protovalidate.ValidationError.

		if errors.As(err, &validationError) {
			var validationErrorResponse []*common.ValidationError = make([]*common.ValidationError, 0)
			for _, violation := range validationError.Violations { //? Di Laravel, kalau validasi gagal, kita bisa ambil $errors->first('field') atau $errors->all()., Di Go, validationError.Violations itu array yang mirip Laravel: $errors->all() atau $errors->toArray()
				//? Di Laravel, framework otomatis ubah array error ke JSON.
				//? Di Go, kita harus manual mapping violation → response struct.
				validationErrorResponse = append(validationErrorResponse, &common.ValidationError{
					Field:   *violation.Proto.Field.Elements[0].FieldName, //?FieldPath → nama field (name, email)
					Message: *violation.Proto.Message,                     //?Message → mirip pesan error di Laravel (The name field is required.).
				})
			}
			return validationErrorResponse, nil
		}
		return nil, err

	}

	return make([]*common.ValidationError, 0), nil //?"validasi sukses, tidak ada error".
	//?analogi:
	//?  $validated = $request->validate([
	//? 	'name' => 'required',
	//? 	'email' => 'required|email',
	//? ]);
	//? Kalau berhasil → $validated berisi data, dan $errors kosong.
}
