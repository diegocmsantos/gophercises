package main

import (
	"fmt"
	"gophercises/sitemap"
)

func main() {
	// 	html := `
	// 	<html>
	// <head>
	//   <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/font-awesome/4.7.0/css/font-awesome.min.css">
	// </head>
	// <body>
	//   <h1>Social stuffs</h1>
	//   <div>
	//     <a href="https://www.twitter.com/joncalhoun">
	//       Check me out on twitter
	//       <i class="fa fa-twitter" aria-hidden="true"></i>
	//     </a>
	//     <a href="https://github.com/gophercises">
	//       Gophercises is on <strong>Github</strong>!
	//     </a>s
	//   </div>
	// </body>
	// </html>
	// 	`
	// r := strings.NewReader(html)
	url := "https://www.calhoun.io/"
	bytes, _ := sitemap.NewSitemap(url)
	fmt.Printf("%+v\n", string(bytes))
}
