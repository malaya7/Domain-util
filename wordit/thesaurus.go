package wordit

type Thesaurus interface {
	Synonyms(term string) ([]string, error)
}
