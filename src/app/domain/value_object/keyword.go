package value_object

import "errors"

const MaxKeywordLength = 144

var OverLengthLimitTwitterKeywordException = errors.New("the length of keyword has exceeded the limit")

type Keyword struct {
	keyword string
}

func NewKeyword(keyword string) (t *Keyword, err error) {
	if len(keyword) > MaxKeywordLength {
		err = OverLengthLimitTwitterKeywordException
	}
	if err != nil {
		return
	}

	t = new(Keyword)
	t.keyword = keyword
	return
}
