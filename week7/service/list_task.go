package service

import(
	"context"
	"fmt"

	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/entity"
	"github.com/JinHyeokOh01/gdg-on-campus-khu-backend/week7/store"
)

type ListTask struct{
	DB store.Queryer
	Repo TaskLister
}

func (l *ListTask) ListTasks(ctx context.Context) (entity.Tasks, error){
	ts, err := l.Repo.ListTasks(ctx, l.DB)
	if err != nil{
		return nil, fmt.Errorf("failed to list: %w", err)
	}
	return ts, nil
}