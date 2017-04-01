package main

import (
	"crypto/md5"
	"fmt"
	"github.com/exponential-decay/httpreserve"
)

var structids []string

// makeHash will creat a MD5 hash for us to use to index our data without
// duplication...
func makeHash(js string) string {
	md5 := md5.New()
	md5.Write([]byte(js))
	return fmt.Sprintf("%x", md5.Sum(nil))
}

// storeStruct allows us to get a different representation of the LinkStats structure
// e.g. as a map we have good flexibility over looping and passing around without
// reglection to iterate through the struct for us.
func storeStruct(ls httpreserve.LinkStats, js string) map[string]interface{} {

	var lmap = make(map[string]interface{})

	id := makeHash(js)
	for _, y := range structids {
		if id == y {
			return lmap
		}
	}

	structids = append(structids, id)

	lmap["id"] = id
	lmap["filename"] = ls.FileName
	lmap["analysis version number"] = ls.AnalysisVersionNumber
	lmap["analysis version text"] = ls.AnalysisVersionText
	lmap["link"] = ls.Link
	lmap["response code"] = ls.ResponseCode
	lmap["response text"] = ls.ResponseText
	lmap["screen shot"] = ls.ScreenShot
	lmap["internet archive latest"] = ls.InternetArchiveLinkLatest
	lmap["internet archive earliest"] = ls.InternetArchiveLinkEarliest
	lmap["internet archive save link"] = ls.InternetArchiveSaveLink
	lmap["internet archive response code"] = ls.InternetArchiveResponseCode
	lmap["internet archive response text"] = ls.InternetArchiveResponseText
	lmap["archived"] = ls.Archived
	lmap["protocol error"] = ls.ProtocolError
	lmap["protocol error"] = ls.ProtocolErrorMessage

	return lmap
}

//for now, for testing...
var linkmap = map[string]string{
	"http://www.bbc.co.uk/news":           "bbc news",
	"http://www.bbc.co.uk/":               "bbc home",
	"http://www.bbc.co.uk/radio":          "bbc radio",
	"http://www.nationalarchives.gov.uk/": "tna",
}