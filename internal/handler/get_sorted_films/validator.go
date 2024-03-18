package get_sorted_films

const (
	titleOrder       = "title"
	ratingOrder      = "rating"
	releaseDateOrder = "release_date"
)

func validateQueryParam(param string) (bool, string) {
	if param != titleOrder && param != ratingOrder && param != releaseDateOrder {
		return false, "invalid order field"
	}

	return true, ""
}
