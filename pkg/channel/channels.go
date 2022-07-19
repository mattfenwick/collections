package channel

import "context"

// TODO decide how to handle three types of channels (in/out, in, out)

func Merge[A any](ctx context.Context, xs []<-chan A) <-chan A {
	combined := make(chan A)
	for _, x := range xs {
		go func(channel <-chan A) {
			for val := range channel {
				select {
				case <-ctx.Done():
					return
				case combined <- val:
				}
			}
		}(x)
	}
	return combined
}

// Split copies events from a single in channel to multiple out channels.
//   So, a broadcast.
func Split[A any](ctx context.Context, in <-chan A, outs []chan<- A) {
	go func() {
		select {
		case <-ctx.Done():
			return
		case val := <-in:
			for _, out := range outs {
				select {
				case out <- val:
				default:
					// if channel is blocked, this event gets dropped on the floor
				}
			}
		}
	}()
}
