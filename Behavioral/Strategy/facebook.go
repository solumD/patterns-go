package strategy

import "fmt"

type FacebookStrategy struct {
}

func (f FacebookStrategy) NewPost(p Post) {
	fmt.Printf("New post in Facebook. Title: %s | Text: %s", p.Title, p.Text)
}