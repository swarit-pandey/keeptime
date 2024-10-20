package convert

import (
	"fmt"
	"strings"
	"time"

	"github.com/swarit-pandey/keeptime/pkg/common"
)

type Convert struct {
	Opts *Options
}

func NewConvert(opts *Options) *Convert {
	return &Convert{
		Opts: opts,
	}
}

func (c *Convert) Start() error {
	fromLoc, err := common.GetLocation(c.Opts.From)
	if err != nil {
		return err
	}

	toLoc, err := common.GetLocation(c.Opts.To)
	if err != nil {
		return nil
	}

	var t time.Time

	if strings.ToLower(c.Opts.Time) == "now" {
		t = time.Now().In(fromLoc)
	} else if strings.ContainsAny(c.Opts.Time, "smdy") {
		t, err = common.ParseRelativeTime(c.Opts.Time, fromLoc)
		if err != nil {
			return err
		}
	} else {
		t, err = common.ParseTimeLoc(c.Opts.Time, fromLoc)
		if err != nil {
			return err
		}
	}

	convertedTime := t.In(toLoc)
	c.formatAndPrint(convertedTime)
	return nil
}

func (c *Convert) formatAndPrint(t time.Time) {
	switch strings.ToLower(c.Opts.Format.Format) {
	case "rfc3339":
		fmt.Println(t.Format(time.RFC3339))
	case "unix":
		fmt.Println(fmt.Sprintf("%d", t.Unix()))
	case "iso8601":
		fmt.Println(t.Format("2006-01-02T15:04:05Z07:00"))
	case "ansic":
		fmt.Println(t.Format(time.ANSIC))
	default:
		fmt.Println(t.Format(time.RFC3339))
	}
}
