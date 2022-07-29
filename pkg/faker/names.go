package faker

var firstNames = []string{
	"Abbey",
	"Abbie",
	"Abby",
	"Abigail",
	"Ada",
	"Adah",
	"Adaline",
	"Addie",
	"Adela",
	"Adelaida",
	"Adelaide",
	"Adele",
	"Adelia",
	"Adelina",
	"Adeline",
	"Adell",
	"Adella",
	"Adelle",
	"Adena",
	"Adina",
	"Adria",
	"Adrian",
	"Adriana",
	"Adriane",
	"Adrianna",
	"Adrianne",
	"Adrien",
}

var lastNames = []string{
	"Abbott",
	"Abernathy",
	"Abshire",
	"Adams",
	"Altenwerth",
	"Anderson",
	"Ankunding",
	"Armstrong",
	"Auer",
	"Aufderhar",
	"Bahringer",
	"Bailey",
	"Balistreri",
	"Barrows",
	"Bartell",
	"Bartoletti",
	"Barton",
	"Bashirian",
	"Batz",
	"Bauch",
	"Baumbach",
	"Bayer",
	"Beahan",
	"Beatty",
}

func FirstName() string {
	return fetchSample(firstNames)
}

func LastName() string {
	return fetchSample(lastNames)
}

func Name() string {
	return FirstName() + " " + LastName()
}
