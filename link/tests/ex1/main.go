package main

import (
	"fmt"
	"gophercises/link"
	"strings"
)

func main() {
	html := `
	<html>
<body>
  <h1>Hello!</h1>
  <a href="/other-page">A link to another page</a>
</body>
</html>
	`
	r := strings.NewReader(html)
	links, err := link.Parse(r)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)

}
