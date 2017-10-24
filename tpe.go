package main

import "path/filepath"

func tpePath(name string) string {
	return filepath.Join(getCurPath(), "./template/"+name+".tpe")
}

func imgPath(name string) string {
	return filepath.Join(getCurPath(), "./image/"+name+".bmp")
}
