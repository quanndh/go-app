package main

import (
	"github.com/quanndh/go-app/public/boostrap"
	"go.uber.org/fx"
)

func main() {
	fx.New(fx.Options(boostrap.All()...)).Run()
}
