package entitiy

type Opinion struct {
	Id      uint64
	Content string
}

func NewOpinion(id uint64, content string) (opinion *Opinion, err error) {
	opinion = new(Opinion)
	opinion.Id = id
	opinion.Content = content
	return
}
