package search

type defaultMatcher struct{} //defaultMather implements the default matcher

func init() { // init registers the default matcher with the program
	var matcher defaultMatcher
	Register("default", matcher)
}

// implements behaviour for matcher
func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
