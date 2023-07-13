package helpers

func hasSameResidenceID(res1, res2 map[string]interface{}) bool {
	return res1["residenceId"] == res2["residenceId"]
}

func hasSameAddress(res1, res2 map[string]interface{}) bool {
	if res2["streetAddress"] == nil || res2["city"] == nil {
		return false
	}

	return (res1["streetAddress"] == res2["streetAddress"] && res1["city"] == res2["city"])
}

func updateResidenceData(existingResidence map[string]interface{}, residence map[string]interface{}) {
	for k, v := range residence {
		existingResidence[k] = v
	}
}

func UpdateResidenceData(existingResidences []map[string]interface{}, residence map[string]interface{}) []map[string]interface{} {
	var updated bool = false // Flag to check if residence has been updated

	// Check if new residence is set as currentResidence
	newResidenceCurrent := false
	if residence["currentResidence"] != nil {
		newResidenceCurrent = residence["currentResidence"].(bool)
	}

	// set currentResidence to false for all existing residences 
	if newResidenceCurrent {
		for _, existingResidence := range existingResidences {
			existingResidence["currentResidence"] = false
		}
	}

	for _, existingResidence := range existingResidences {
		if updated {
			break
		}

		if hasSameResidenceID(existingResidence, residence) {
			updateResidenceData(existingResidence, residence)

			updated = true
		} else if hasSameAddress(existingResidence, residence) {
			updateResidenceData(existingResidence, residence)

			updated = true
		}

	}

	// If residence has not been updated, append it to the slice
	if !updated {
		existingResidences = append(existingResidences, residence)
	}

	return existingResidences
}
