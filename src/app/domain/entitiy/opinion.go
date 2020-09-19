package entitiy

import (
	"regexp"
	"strings"
)

type Opinion struct {
	Id      uint64
	Content string
}

func NewOpinion(id uint64, content string) (opinion *Opinion, err error) {
	opinion = new(Opinion)
	opinion.Id = id
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "RT", "", -1)
	r := regexp.MustCompile(`@\w+:*\s*`)
	content = r.ReplaceAllString(content, "")
	r = regexp.MustCompile(`https?://[\w!?/+\-_~;.,*&@#$%()'[\]]+`)
	content = r.ReplaceAllString(content, "")

	opinion.Content = content
	return
}
