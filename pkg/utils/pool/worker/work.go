package worker

import "context"

func (n *Worker[T]) Work(ctx context.Context) {
	// Implementation specific to the worker type
	result, err := n.work(ctx)
	if err != nil {
		n.Publisher.NotifyErr(err)
		return
	}
	n.Publisher.Notify(result)
}
