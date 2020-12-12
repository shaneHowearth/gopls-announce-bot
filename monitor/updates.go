// Package monitor -
package monitor

// Store -
type Store interface {
	GetReleases() (proverb, translation, explanation string, err error)
	// GetLatestCommit() (placename, translation, explanation string, err error)
}

/*
func isMorning() bool {
	now := time.Now()
	newLayout := "15:04"
	check, _ := time.Parse(newLayout, strconv.Itoa(now.Hour())+":"+strconv.Itoa(now.Minute()))
	start, _ := time.Parse(newLayout, "23:59")
	end, _ := time.Parse(newLayout, "11:59")
	start, end = start.UTC(), end.UTC()
	if start.After(end) {
		start, end = end, start

	}
	check = check.Local()
	return !check.Before(start) && !check.After(end)

}

// GetRelease -
func GetRelease(set string, s Store) (content, translation, explanation string, err error) {
	switch set {
	case "proverb":
		return s.GetRandomProverb()
	case "placename":
		return s.GetRandomPlacename()
	default:
		return "", "", "", fmt.Errorf("%s does not exist as an option", set)
	}
}
*/
