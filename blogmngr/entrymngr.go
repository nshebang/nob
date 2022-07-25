package blogmngr

import (
	"os"
	"os/exec"
	"io/ioutil"
	"runtime"
	"fmt"
	"sort"
	"strings"
	"time"
	"html"
	"path/filepath"
)

func CreateDraftAndEdit(title string) bool {
	editor := os.Getenv("EDITOR")
	entry := NewEntry(title)

	fpath := fmt.Sprintf("nob/drafts/%s.md", toFilename(title))

	if _, err := os.Stat(fpath); !os.IsNotExist(err) {
		return false
	}

	f, _ := os.OpenFile(fpath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0600)
	f.WriteString(entry.Draft())
	f.Close()

	if runtime.GOOS == "windows" {
		fmt.Println("Edit the new entry using your favorite text editor.")
		return true
	}
	if editor == "" {
		fmt.Println("Set the environment var EDITOR to use a custom text editor")
		editor = "nano"
	}
	cmd := exec.Command(editor, fpath)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()
	return true
}

func BuildStatic(autodelete bool) {
	var entriesUnsorted = make(map[int64]Entry)
	loadEntries(entriesUnsorted)
	entries := make([]Entry, 0)
	sortEntries(entriesUnsorted, &entries)
	overwriteStaticFiles(&entries)

	if autodelete {
		autoDelete(&entries)
	}
}

func autoDelete(entries *[]Entry) {
	var entriesHTML = make([]string, 0)
	var titles = make([]string, 0)
	loadHTMLEntries(&entriesHTML)
	
	for _, entry := range (*entries) {
		t := fmt.Sprintf("%s.html", toFilename(entry.title))
		titles = append(titles, t)
	}
	
	for _, fname := range entriesHTML {
		exists := false
		for _, t := range titles {
			if t == fname {
				exists = true
			}
		}
		if !exists {
			os.Remove(fname)
		}
	}
}

func toFilename(str string) string {
	fname := strings.ToLower(str)
	fname = strings.Replace(fname, " ", "_", -1)
	fname = strings.Replace(fname, "'", "", -1)
	fname = strings.Replace(fname, "\"", "", -1)
	fname = strings.Replace(fname, "#", "", -1)
	fname = strings.Replace(fname, "/", "_", -1)
	return fname
}

func loadEntries(emap map[int64]Entry) {
	drafts, _ := ioutil.ReadDir("nob/drafts/")

	for _, file := range drafts {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}
		efile := fmt.Sprintf("nob/drafts/%s", file.Name())
		entry := NewEntryFromFile(efile)
		time, _ := time.Parse(
			"Mon, 01 Jan 2006 15:04:05 -0700",
			entry.dateRfc2822)
		timestamp := time.Unix()
		emap[timestamp] = (*entry)
	}
}

func loadHTMLEntries(entries *[]string) {
	flist, _ := ioutil.ReadDir(".")
	
	for _, file := range flist {
		fname := file.Name()
		if filepath.Ext(fname) != ".html" || fname == "index.html" {
			continue
		}
		(*entries) = append((*entries), fname)
	}
}

func sortEntries(emap map[int64]Entry, entries *[]Entry) {
	times := make([]int64, 0)
	for k := range emap {
		times = append(times, k)
	}
	
	sort.Slice(times, func(i, j int) bool { return times[i] < times[j]})
	
	for i := len(times) - 1; i >= 0; i-- {
		(*entries) = append((*entries), emap[times[i]])
	}
}

func overwriteStaticFiles(entries *[]Entry) {
	var indexHTML string
	var rssXML string

	var entriesHTML string
	var entriesXML string
	
	var currentHTML string

	for _, entry := range (*entries) {
		file := fmt.Sprintf("%s.html", toFilename(entry.title))
		ehtml := entry.ToHTML()

		entriesHTML += fmt.Sprintf(EntryLi, file, entry.title, entry.date)	
		entriesHTML += "    "

		entriesXML += FmtTemplate(RssArticle, P{
			"title": entry.title,
			"link": fmt.Sprintf("%s/%s", GetWebsiteURL(), file),
			"desc": html.UnescapeString(ehtml),
		})

		currentHTML = ReadTemplateFile("post.html")
		currentHTML = FmtTemplate(currentHTML, P{
			"title": entry.title,
			"date": entry.date,
			"content": ehtml,
		})
		WriteToStatic(file, currentHTML)
	}

	indexHTML = ReadTemplateFile("index.html")
	indexHTML = FmtTemplate(indexHTML, P{"entries": entriesHTML})

	rssXML = ReadTemplateFile("rss.xml")
	rssXML = FmtTemplate(rssXML, P{"link": GetWebsiteURL(), "entries": entriesXML})

	WriteToStatic("index.html", indexHTML)
	WriteToStatic("rss.xml", rssXML)
}

