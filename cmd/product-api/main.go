package main

import "product-api/app"

func main() {
	app.New().Start("0.0.0.0", 3000)
}
