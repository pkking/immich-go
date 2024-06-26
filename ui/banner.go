package ui

import (
	"fmt"
	"strings"
)

type Banner struct {
	b                     []string
	version, commit, date string
}

// Banner Ascii art
// Generator : http://patorjk.com/software/taag-v1/
// Font: Three point

func NewBanner(version, commit, date string) Banner {
	return Banner{
		b: []string{
			". _ _  _ _ . _|_     _  _ ",
			"|| | || | ||(_| | ─ (_|(_)",
			"                     _)   ",
		},
		version: version,
		commit:  commit,
		date:    date,
	}
}

// String generate a string with new lines and place the given text on the latest line
func (b Banner) String() string {
	const lenVersion = 20
	var text string
	if b.version != "" {
		text = fmt.Sprintf("v %s", b.version)
	}
	sb := strings.Builder{}
	for i := range b.b {
		if i == len(b.b)-1 && text != "" {
			if len(text) >= lenVersion {
				text = text[:lenVersion]
			}
			sb.WriteString(b.b[i][:lenVersion-len(text)] + text + b.b[i][lenVersion:])
		} else {
			sb.WriteString(b.b[i])
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}
