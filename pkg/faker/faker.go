package faker

import "math/rand"

// fetchSample returns a random item from the given slice.
func fetchSample[v int | string](s []v) v {
	return s[rand.Intn(len(s))]
}
