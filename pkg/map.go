package pkg

func MapMap[A comparable, B, C any](f F1[B, C], kvs map[A]B) map[A]C {
	out := map[A]C{}
	for k, v := range kvs {
		out[k] = f(v)
	}
	return out
}
