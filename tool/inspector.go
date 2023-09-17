package tool

import "strings"

var illegalWords = []string{
	"习进", "反gong", "操", "NM", "nm", "sb", "SB",
	"<script","select*","SELECT*","ORDERBY","orderby",
	"andif","ANDIF","SLEEP()","sleep()","=","+","\\",",","'",
}

func WordsInspect(content string) bool {
	if content == "" {
		return false
	}
	for _, v := range illegalWords {
		if strings.Contains(content, v) {
			return false
		} 
	}
	return true
}