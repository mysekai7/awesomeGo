package main

import (
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

/**
command:
go run main.go dsd --verbose -i image
out:
true, dsd, image
*/

var (
	verbose = kingpin.Flag("verbose", "Verbose mode.").Short('v').Bool()
	//postImage = kingpin.Flag("image", "Image to post.").Short('i').String()
	name = kingpin.Arg("cert", "Machine cert.").String()
	//items     = kingpin.Flag("items", "Product item").StringMap()
	item     = kingpin.Flag("item", "products items").Short('i').StringMap()
	products = kingpin.Flag("products", "products items").Short('p').Strings()
)

func main() {

	kingpin.Parse()
	fmt.Printf("%v, %s\n", *verbose, *name)
	fmt.Printf("item: %+v\n", *item)
	fmt.Printf("products: %+v\n", *products)
}
