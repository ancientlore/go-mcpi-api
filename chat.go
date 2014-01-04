package mcpiapi

import (
	"fmt"
	"strings"
)

type Chat object

func (obj Chat) Post(text string) error {
	s := fmt.Sprintf("chat.post(\"%s\")", strings.Replace(text, "\"", "\\\"", -1))
	return object(obj).send(s)
}
