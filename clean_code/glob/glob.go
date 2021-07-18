package glob

import "regexp"

func Match(p, s string) (bool, error) {
	// pattern := translate(p)
	re, err := regexp.Compile(p)
	if err != nil {
		return false, err
	}
	isMatch := re.MatchString(s)
	return isMatch, nil
}
