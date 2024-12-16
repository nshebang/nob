package blogmngr

import (
	"os"
	"os/exec"
	"io/ioutil"
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

	fpath := fmt.Sprintf(".nob/entries/%s.md", toFilename(title))

	if _, err := os.Stat(fpath); !os.IsNotExist(err) {
		i := 0
		for {
			ext := filepath.Ext(fpath)
			base := fpath[:len(fpath) - len(ext)]
			i++
			testPath := fmt.Sprintf("%s-%d%s", base, i, ext)
			
			_, err := os.Stat(testPath)
			if os.IsNotExist(err) {
				fpath = testPath
				break
			}
		}
	}

	f, _ := os.OpenFile(fpath, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0600)
	f.WriteString(entry.Markdown())
	f.Close()

	if editor == "" {
		fmt.Println("Set the environment var $EDITOR to use a custom text editor")
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
		fnameWithExt := filepath.Base(entry.path)
		fname := strings.TrimSuffix(fnameWithExt, filepath.Ext(fnameWithExt))
		title := fmt.Sprintf("%s.html", fname)
		titles = append(titles, title)
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
	fname = strings.Replace(fname, " ", "-", -1)
	fname = strings.Replace(fname, "'", "", -1)
	fname = strings.Replace(fname, "\"", "", -1)
	fname = strings.Replace(fname, "#", "", -1)
	fname = strings.Replace(fname, "/", "-", -1)
	fname = strings.Replace(fname, "\\", "-", -1)
	return fname
}

func loadEntries(entries map[int64]Entry) {
	entryList, _ := ioutil.ReadDir(".nob/entries/")

	for _, file := range entryList {
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}
		entryFile := fmt.Sprintf(".nob/entries/%s", file.Name())
		entry := NewEntryFromFile(entryFile)

		if entry.Draft() {
			continue
		}

		time, _ := time.Parse(
			"Mon, 01 Jan 2006 15:04:05 -0700",
			entry.dateRfc2822)
		timestamp := time.Unix()
		entries[timestamp] = (*entry)
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

func sortEntries(entryTimestamps map[int64]Entry, entries *[]Entry) {
	timestamps := make([]int64, 0, len(entryTimestamps))
	for k := range entryTimestamps {
		timestamps = append(timestamps, k)
	}
	
	sort.Slice(timestamps, func(i, j int) bool { return timestamps[i] < timestamps[j]})
	
	for i := len(timestamps) - 1; i >= 0; i-- {
		(*entries) = append((*entries), entryTimestamps[timestamps[i]])
	}
}

func overwriteStaticFiles(entries *[]Entry) {
	var indexHTML string
	var rssXML string

	var entriesHTML string
	var entriesXML string
	
	var currentHTML string

	for _, entry := range (*entries) {
		fnameWithExt := filepath.Base(entry.path)
		fname := strings.TrimSuffix(fnameWithExt, filepath.Ext(fnameWithExt))
		entryPath := fmt.Sprintf("%s.html", fname)
		
		ehtml := entry.ToHTML()

		entriesHTML += fmt.Sprintf(EntryLi, entryPath, entry.title, entry.date)	
		entriesHTML += "    "

		entriesXML += FmtTemplate(RssArticle, P{
			"title": entry.title,
			"link": fmt.Sprintf("%s/%s", GetWebsiteURL(), entryPath),
			"desc": html.UnescapeString(ehtml),
		})

		currentHTML = ReadTemplateFile("post.html")
		currentHTML = FmtTemplate(currentHTML, P{
			"title": entry.title,
			"date": entry.date,
			"content": ehtml,
		})
		WriteToStatic(entryPath, currentHTML)
	}

	indexHTML = ReadTemplateFile("index.html")
	indexHTML = FmtTemplate(indexHTML, P{"entries": entriesHTML})

	rssXML = ReadTemplateFile("rss.xml")
	rssXML = FmtTemplate(rssXML, P{"link": GetWebsiteURL(), "entries": entriesXML})

	WriteToStatic("index.html", indexHTML)
	WriteToStatic("rss.xml", rssXML)
}

