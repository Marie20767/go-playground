package arrays

func PushLastUuidToFront(ids []string) []string {
	// assume that all ids are unique
	if len(ids) == 0 {
		return ids
	}

	lastId := ids[len(ids)-1]
	idsWithoutLastId := ids[:len(ids)-1]
	reorderedIds := append([]string{lastId}, idsWithoutLastId...)

	return reorderedIds
}

func PushLastUuidToFrontInPlace(ids []string) []string {
	// assume that all ids are unique
	if len(ids) == 0 {
		return ids
	}

	originalIds := make([]string, len(ids))
	copy(originalIds, ids)

	for i := range ids {
		if i == len(ids)-1 {
			ids[0] = originalIds[i]
		} else {
			ids[i+1] = originalIds[i]
		}
	}

	return ids
}
