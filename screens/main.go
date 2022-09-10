package screens

import "net/url"

var GIT_URL, YT_URL, TWT_URL *url.URL

func initURLs() error {
	var err error

	GIT_URL, err = url.Parse("https://github.com/MajestikButter")
	if err != nil {
		return err
	}

	YT_URL, err = url.Parse("https://www.youtube.com/channel/UCnl-jxpL_DMSicqCoGwCA2g")
	if err != nil {
		return err
	}

	TWT_URL, err = url.Parse("https://twitter.com/MajestikButter1")
	if err != nil {
		return err
	}

	return nil
}

func init() {
	initURLs()
}
