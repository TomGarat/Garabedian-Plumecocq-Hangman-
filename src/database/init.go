package database

var (
	m = Mots{}
)

type Mots struct {
	Word []string
}

func (m *Mots) Init() {
	m.Word = []string{}
}
