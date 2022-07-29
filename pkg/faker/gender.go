package faker

var gender = []string{"male", "female", "non-binary", "agender", "androgyne", "bi-gender", "pan-gender"}

func Gender() string {
	return fetchSample(gender)
}
