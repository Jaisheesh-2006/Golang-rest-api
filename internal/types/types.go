package types

type Student struct {
	Id    int64 `json:"id" `
	Name  string `json:"name" validate:"required,min=2,max=100"`
	Age   int `json:"age" validate:"required,min=1,max=150"`
	Email string `json:"email" validate:"required,email"`
}
