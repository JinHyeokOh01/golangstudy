package service

import(
	"context"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/entity"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/store"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . TaskAdder TaskLister UserRegister
type TaskAdder interface{
	AddTask(ctx context.Context, db store.Execer, t *entity.Task) error
}

type TaskLister interface{
	ListTasks(ctx context.Context, db store.Queryer) (entity.Tasks, error)
}

type UserRegister interface{
	RegisterUser(ctx context.Context, db store.Execer, u *entity.User) error
}