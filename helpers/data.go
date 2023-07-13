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

	if !updated {
		existingResidences = append(existingResidences, residence)
	}

	return existingResidences
}
