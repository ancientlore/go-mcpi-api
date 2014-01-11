package mcpiapi

import (
	"fmt"
	"strings"
)

// Chat provides methods to post messages to the user.
type Chat object

// Post displays the given text to the user.
func (obj Chat) Post(text string) error {
	s := fmt.Sprintf("chat.post(\"%s\")", strings.Replace(text, "\"", "\\\"", -1))
	return object(obj).send(s)
}
