package checks

// GoCyclo is the check for the go cyclo command
type GoCyclo struct {
	Dir       string
	Filenames []string
}

// Name returns the name of the display name of the command
func (g GoCyclo) Name() string {
	return "gocyclo"
}

// Weight returns the weight this check has in the overall average
func (g GoCyclo) Weight() float64 {
	return .10
}

// Percentage returns the percentage of .go files that pass gofmt
func (g GoCyclo) Percentage() (float64, []FileSummary, error) {
	return GoTool(g.Dir, g.Filenames, []string{"gometalinter", "--deadline=180s", "--disable-all", "--enable=gocyclo", "--cyclo-over=15"})
}

// Description returns the description of GoCyclo
func (g GoCyclo) Description() string {
	return "Gocyclo calculates cyclomatic complexities of functions in Go source code."
}
