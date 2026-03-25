package pages

import (
	"context"
	"fmt"
	"kitt/kitt/repository"
	"time"
)

type SyncWorker struct {
	Pages repository.Repository[Page, int64]
}

func (w *SyncWorker) Id() string {
	return "admin.pages.sync"
}

func (w *SyncWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			// sync pages or fetch remote data
			fmt.Println("TICK FROM WORKER.")
		}
	}
}
