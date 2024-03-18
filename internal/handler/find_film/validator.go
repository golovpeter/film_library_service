package find_film

const (
	titleSearchField     = "title"
	actorNameSearchField = "actor"
)

func validateQueryParam(param string) (bool, string) {
	if param == "" {
		return false, "search field must not be empty"
	}

	if param != titleSearchField && param != actorNameSearchField {
		return false, "invalid search field"
	}

	return true, ""
}
