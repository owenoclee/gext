package models

import (
	"fmt"
	"html/template"
	"regexp"
)

type Post struct {
	Id        uint32
	ReplyTo   uint32
	Board     string
	Subject   string
	Body      string
	CreatedAt uint32
}

func (p Post) ThreadID() uint32 {
	if p.ReplyTo == 0 {
		return p.Id
	}
	return p.ReplyTo
}

func (p Post) BodySafeHTML() template.HTML {
	p.Body = template.HTMLEscapeString(p.Body)
	p = replaceSameThreadReplies(p)
	p = replaceDifferentThreadReply(p)
	p = replaceDifferentBoardReply(p)
	return template.HTML(p.Body)
}

var sameThreadReply = regexp.MustCompile(`(\A|\s)&gt;&gt;(\d{1,10})(\s|\z)`)

func replaceSameThreadReplies(p Post) Post {
	for i := 0; i < 2; i++ {
		p.Body = ReplaceAllStringSubmatchFunc(sameThreadReply, p.Body, func(subs []string) string {
			return fmt.Sprintf(
				"%v<a href=\"/%v/thread/%v#%v\">&gt;&gt;%v</a>%v",
				subs[1],
				p.Board,
				p.ThreadID(),
				subs[2],
				subs[2],
				subs[3],
			)
		})
	}
	return p
}

var differentThreadReply = regexp.MustCompile(`(\A|\s)&gt;&gt;/(\d{1,10})/(\d{1,10})(\s|\z)`)

func replaceDifferentThreadReply(p Post) Post {
	for i := 0; i < 2; i++ {
		p.Body = ReplaceAllStringSubmatchFunc(differentThreadReply, p.Body, func(subs []string) string {
			return fmt.Sprintf(
				"%v<a href=\"/%v/thread/%v#%v\">&gt;&gt;/%v/%v</a>%v",
				subs[1],
				p.Board,
				subs[2],
				subs[3],
				subs[2],
				subs[3],
				subs[4],
			)
		})
	}
	return p
}

var differentBoardReply = regexp.MustCompile(`(\A|\s)&gt;&gt;&gt;/([a-z]{1,16})/(\d{1,10})/(\d{1,10})(\s|\z)`)

func replaceDifferentBoardReply(p Post) Post {
	for i := 0; i < 2; i++ {
		p.Body = ReplaceAllStringSubmatchFunc(differentBoardReply, p.Body, func(subs []string) string {
			return fmt.Sprintf(
				"%v<a href=\"/%v/thread/%v#%v\">&gt;&gt;&gt;/%v/%v/%v</a>%v",
				subs[1],
				subs[2],
				subs[3],
				subs[4],
				subs[2],
				subs[3],
				subs[4],
				subs[5],
			)
		})
	}
	return p
}

func ReplaceAllStringSubmatchFunc(re *regexp.Regexp, str string, repl func([]string) string) string {
	result := ""
	lastIndex := 0

	for _, v := range re.FindAllSubmatchIndex([]byte(str), -1) {
		groups := []string{}
		for i := 0; i < len(v); i += 2 {
			groups = append(groups, str[v[i]:v[i+1]])
		}

		result += str[lastIndex:v[0]] + repl(groups)
		lastIndex = v[1]
	}

	return result + str[lastIndex:]
}
