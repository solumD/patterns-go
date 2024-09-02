package strategy

import "fmt"

type InstagramStrategy struct {
}

func (i InstagramStrategy) NewPost(p Post) {
	fmt.Printf("New post in Instagram. Title: %s | Text: %s\n", p.Title, p.Text)
}