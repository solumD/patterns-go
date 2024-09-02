package strategy

// Интерфейс стратегии
type Strategy interface {
    NewPost(Post)
}

type Post struct {
    Title string
    Text string
}