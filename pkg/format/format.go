package format

import (
	"fmt"
	"strings"
	"time"

	"github.com/swarit-pandey/keeptime/pkg/common"
)

type Format struct {
	Opts *Options
}

func NewFormat(opts *Options) *Format {
	return &Format{
		Opts: opts,
	}
}

func (f *Format) Start() error {
	loc, err := common.GetLocation(f.Opts.System)
	if err != nil {
		return err
	}

	var t time.Time

	if strings.ToLower(f.Opts.Time) == "now" {
		t = time.Now().In(loc)
	} else if strings.ContainsAny(f.Opts.Time, "smdy") {
		t, err = common.ParseRelativeTime(f.Opts.Time, loc)
	} else {
		t, err = common.ParseTimeLoc(f.Opts.Time, loc)
		if err != nil {
			return err
		}
	}

	f.formatAndPrint(t)
	return nil
}

func (f *Format) formatAndPrint(t time.Time) {
	switch strings.ToLower(f.Opts.Format) {
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
