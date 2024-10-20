package common

// Default Layout
const (
	DefaultLayout = "20060102150405" // YYYYMMDDHHMMSS
)

// TimeZones
const (
	IST   = "Asia/Kolkata"
	EST   = "America/New_York"
	PST   = "America/Los_Angeles"
	CET   = "Europe/Berlin"
	UTC   = "UTC"
	GMT   = "Europe/London"
	JST   = "Asia/Tokyo"
	AEDT  = "Australia/Sydney"
	Local = "Local"
)

// TimeZones maps tz abbreviation to actual location
var TimeZones = map[string]string{
	"UTC":   UTC,
	"IST":   IST,
	"EST":   EST,
	"PST":   PST,
	"CET":   CET,
	"GMT":   GMT,
	"JST":   JST,
	"AEDT":  AEDT,
	"LOCAL": Local,
}
