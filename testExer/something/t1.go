package main

import (
	"fmt"
	"path/filepath"
)

type paths struct {
	name    string
	dirPath bool
	exist   bool
}

func (p *paths) initP(s string) error {
	p.name = s
	fi, fiErr := os.Lstat(s)
	if fiErr != nil {
		lg.Fatalln(fiErr)
	}
	if fi == nil {
		p.exist = false
		if strings.HasSuffix(s, "/") {
			p.dirPath = true
			return nil
		}
		p.dirPath = false
		return nil
	}
	p.exist = true
	if fi.IsDir() {
		p.dirPath = true
		return nil
	}
	p.dirPath = false
	return nil
}

var dP paths

func check(d string) bool {
	// sP := paths{}
	// dP = paths{}
	// sP.initP(s)
	dP.initP(d)
	if dP.dirPath && dP.exist {
		return true
	}
	if (!dP.dirPath) && (!dP.exist) {
		return true
	}
	return false
}
