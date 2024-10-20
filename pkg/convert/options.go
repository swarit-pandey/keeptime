package convert

type Options struct {
	From string
	To   string
	Time string

	Format Format
}

// This is prolly redundant

type Format struct {
	Format string
}
