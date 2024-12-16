package blogmngr

import (
	"fmt"
	"time"
	"io/ioutil"
	"bufio"
	"strings"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

type Entry struct {
	title		string
	date		string
	dateRfc2822	string
	contentMd	string
	draft		bool
}

const mdTemplate = `---
TITLE: %s
DATE: %s
DATE RFC2822: %s
DRAFT: false
---
%s
`

func NewEntry(newTitle string) *Entry {
	now := time.Now().UTC()
	newDate := now.Format("01-02-2006")
	newRfc2822 := now.Format("Mon, 01 Jan 2006 15:04:05 -0700")

	return &Entry{
		title: newTitle,
		date: newDate,
		dateRfc2822: newRfc2822,
		contentMd: "",
		draft: false,
	}
}

func NewEntryFromFile(fpath string) *Entry {
	f, _ := ioutil.ReadFile(fpath)
	raw := string(f)

	title, date, rfc2822, draft := parseTitleBlock(raw)
	
	mdStart := strings.Index(raw, "---\n")
	mdEnd := strings.Index(raw[mdStart + 4:], "---\n")

	content := string(raw[mdStart + 4 + mdEnd + 4:])

	return &Entry{
		title: title,
		date: date,
		dateRfc2822: rfc2822,
		contentMd: content,
		draft: draft,
	} 
}

func (e *Entry) Draft() bool {
	return e.draft
}

func (e *Entry) Markdown() string {
	return fmt.Sprintf(mdTemplate, e.title, e.date, e.dateRfc2822, e.contentMd)
}

func (e *Entry) ToHTML() string {
	ext := parser.CommonExtensions | parser.AutoHeadingIDs | parser.SuperSubscript
	mdparser := parser.NewWithExtensions(ext)
	post := e.contentMd

	return string(markdown.ToHTML([]byte(post), mdparser, nil))
}

func parseTitleBlock(raw string) (string, string, string, bool) {
	var title string
	var date string
	var rfc2822 string
	draft := false

	scanner := bufio.NewScanner(strings.NewReader(raw))
	scanner.Split(bufio.ScanLines)
	l := 0

	for scanner.Scan() {
		line := scanner.Text()
		first := strings.Index(line, ": ")

		if l > 5 {
			break
		}
		if strings.Count(line, "TITLE: ") > 0 {
			title = line[first + 2:]
		}
		if strings.Count(line, "DATE: ") > 0 {
			date = line[first + 2:]
		}
		if strings.Count(line, "DATE RFC2822: ") > 0 {
			rfc2822 = line[first + 2:]
		}
		if strings.Count(line, "DRAFT: ") > 0 {
			rawDraft := line[first + 2:]
			if rawDraft == "true" {
				draft = true
			}
		}
		l++
	}
	return title, date, rfc2822, draft
}

