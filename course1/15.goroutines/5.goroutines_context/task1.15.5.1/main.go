package main

import (
	"context"
	"time"
)

func main() {
	var res string
	res = contextWithDeadline(context.Background(), 1*time.Second, 2*time.Second)
	println(res)
	res = contextWithDeadline(context.Background(), 2*time.Second, 1*time.Second)
	println(res)
	/* Output:
	context deadline exceeded
	time after exceeded
	*/
}

func contextWithDeadline(ctx context.Context, contextDeadline time.Duration, timeAfter time.Duration) string {

	// Place your code here?
	/*
		var t time.Time
		t = time.Now()
		tr := t.Add(contextDeadline)
		tv := t.Add(timeAfter)

	*/
	var t time.Time
	t = time.Now()
	tr := t.Add(contextDeadline)
	ctx, cancel := context.WithDeadline(ctx, tr)
	defer cancel()
	select {
	case <-ctx.Done():
		return "context deadline exceeded"
	case <-time.After(timeAfter):
		return "time after exceeded"
	}
}
