package psdSiteParser

import (
	"regexp"
	"strconv"
)

func Parse(site []byte) Profiles {
	// https://stackoverflow.com/a/1732454/4162386
	var profiles Profiles
	var regexExp = regexp.MustCompile(`(?ms)<div id="profil-(?P<id>\d+)".*?row"><div class="cell rank">.*?<a href="(?P<url>[^"]+)" title="Link zum Profil von (?P<name>[^"]+)">.*?<span class="profile-txt-stimmen">(?P<votes>\d+)<\/span>`)

	matches := regexExp.FindAllStringSubmatch(string(site), -1)

	for _, match := range matches {
		groups := make(map[string]string)
		for i, m := range match {
			groups[regexExp.SubexpNames()[i]] = m
		}

		votes, err := strconv.Atoi(groups["votes"])
		if err != nil {
			continue
		}
		profiles = append(profiles, Profile{Name: groups["name"], Id: groups["id"], Votes: votes})
	}

	return profiles
}
