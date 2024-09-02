package strategy

func StrategyExample() {
    strategies := []Strategy{
        &InstagramStrategy{},
        &FacebookStrategy{},
    }

    blog := &Blog{}
    p := Post{
        Title: "Golang",
        Text: "Best language in history",
    }
    for _, str := range strategies {
        blog.SetStrategy(str)
        blog.strategy.NewPost(p)
    }
}