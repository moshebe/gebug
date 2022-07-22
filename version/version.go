package version

var (
	Version  string
	Revision string
	Branch   string
)

func Name() string {
	if Version != "" {
		return Version
	}
	v := ""
	if Branch != "" {
		v = Branch + "-"
	}

	if Revision != "" {
		return v + Revision
	}

	return "development"
}
