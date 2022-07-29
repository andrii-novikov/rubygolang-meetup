package greeter

type Printable interface {
	Printf(format string, a ...any) (n int, err error)
}

func SayHello(printer Printable, name string, gender string) {
	printer.Printf("Hello %s. Are you %s?\n", name, gender)
}
