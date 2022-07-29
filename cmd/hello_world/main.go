package main

import (
	"github.com/andy1341/rubygolang-meetup/internal/greeter"
	"github.com/andy1341/rubygolang-meetup/pkg/faker"
	"github.com/fatih/color"
)

func main() {
	printer := color.New(color.FgRed).Add(color.Bold)

	greeter.SayHello(printer, faker.Name(), faker.Gender())
}
