package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
)

func router(w http.ResponseWriter, r *http.Request) {
	var expr = regexp.MustCompile("/artist/([0-9]+)")
	var assetExp = regexp.MustCompile("/templates/(js|css)/([^/]+).(js|css)$")

	switch {
	case r.URL.Path == "/":
		home(w, r)
	case r.URL.Path == "/all":
		all(w, r)
	case expr.MatchString(r.URL.Path): // if url match with regexp
		m := expr.FindStringSubmatch(r.URL.Path)
		id, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal(err)
		}
		artist(w, r, id)
	case assetExp.MatchString(r.URL.Path):
		loadAsset(w, r, r.URL.Path)
	default:
		http.NotFound(w, r)
	}
}
