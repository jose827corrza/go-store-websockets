package dtos

type TaskRequest struct {
	Title       string `validate:"required"`
	Description string
	Priority    int `validate:"required,oneof='1' '2' '3'"`
}

type TaskUpdate struct {
	Title       string `validate:"required"`
	Description string
	Priority    int `validate:"required,oneof='1' '2' '3'"`
	IsComplete  bool
}

type SubTask struct {
	Name      string `validate:"required"`
	Completed bool
}
