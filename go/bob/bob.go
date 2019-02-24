// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "regexp"
)

const (
    sure = "Sure."
    calm = "Calm down, I know what I'm doing!"
    chill = "Whoa, chill out!"
    whatever = "Whatever."
    fine = "Fine. Be that way!"
)

var reHasAlphaChars, _ = regexp.Compile(`[A-Za-z]`)

func Hey(remark string) string {
    remark = strings.Trim(remark, "\t\n\r ")

    switch {
    case remark == "":
        return fine
    case strings.ToUpper(remark) == remark && !strings.HasSuffix(remark, "?") && reHasAlphaChars.MatchString(remark):
        return chill
    case strings.HasSuffix(remark, "?") && strings.ToUpper(remark) == remark && reHasAlphaChars.MatchString(remark):
        return calm
    case strings.HasSuffix(remark, "?"):
        return sure
    default:
        return whatever
    }
}
