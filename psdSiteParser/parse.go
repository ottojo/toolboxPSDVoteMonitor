package psdSiteParser

import (
	"regexp"
	"strconv"
	"net/url"
)

func Parse(site []byte) Profiles {
	var profiles Profiles
	var regexExp = regexp.MustCompile(`<a class="archive-titel" href="https://www\.psd-miteinander-leben\.de/profile/([^/]+)/" title="[^"]+">([^"]+)</a>.+<span class="profile-txt-votes">Stimmen</span><span class="profile-txt-stimmen">(\d+)</span>`)

	matches := regexExp.FindAllStringSubmatch(string(site), -1)

	for _, match := range matches {
		if len(match) < 4 {
			continue
		}
		votes, err := strconv.Atoi(match[3])
		if err != nil {
			continue
		}
		id, err := url.QueryUnescape(match[1])
		if err != nil{
			continue
		}
		profiles = append(profiles, Profile{Name: match[2], Id: id, Votes: votes})
	}

	return profiles
}
