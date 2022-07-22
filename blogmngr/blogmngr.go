package blogmngr

import (
	"os"
	"fmt"
)

func CreateBlog(dirname string) bool {
	cwd, _ := os.Getwd()
	blogdir := fmt.Sprintf("%s/%s", cwd, dirname)
	nobdir := fmt.Sprintf("%s/nob", blogdir)
	dirs := [2]string{"drafts", "templates"}

	if _, err := os.Stat(blogdir); !os.IsNotExist(err) {
		return false
	}

	os.MkdirAll(nobdir, 0700)
	for _, dir := range dirs {
		path := fmt.Sprintf("%s/%s", nobdir, dir)
		os.Mkdir(path, 0700)
		if dir == "templates" {
			CreateTemplates(path)
		}
	}
	return true
}

func IsBlog() bool {
	cwd, _ := os.Getwd()
	nobdir := fmt.Sprintf("%s/nob", cwd)

	fi, err := os.Stat(nobdir)
	return !os.IsNotExist(err) && fi.IsDir()
}

