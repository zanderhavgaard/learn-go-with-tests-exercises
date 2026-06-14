package blogposts

import (
	"log"
	"os"

	blogposts "github.com/zanderhavgaard/learn-go-with-tests-exercises/reading-files"
)

func main() {
	posts, err := blogposts.NewPostsFromFS(os.DirFS("posts"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(posts)
}
