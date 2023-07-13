package helpers

import (
	"reflect"
	"testing"
)

func TestHasSameResidenceID(t *testing.T) {
	res1 := map[string]interface{}{"residenceId": "123"}
	res2 := map[string]interface{}{"residenceId": "123"}
	if !hasSameResidenceID(res1, res2) {
		t.Errorf("hasSameResidenceID failed when it should have passed")
	}

	res3 := map[string]interface{}{"residenceId": "123"}
	res4 := map[string]interface{}{"residenceId": "456"}
	if hasSameResidenceID(res3, res4) {
		t.Errorf("hasSameResidenceID passed when it should have failed")
	}
}

func TestHasSameAddress(t *testing.T) {
	res1 := map[string]interface{}{"streetAddress": "123 Street", "city": "New York"}
	res2 := map[string]interface{}{"streetAddress": "123 Street", "city": "New York"}
	if !hasSameAddress(res1, res2) {
		t.Errorf("hasSameAddress failed when it should have passed")
	}

	res3 := map[string]interface{}{"streetAddress": "123 Street", "city": "New York"}
	res4 := map[string]interface{}{"streetAddress": "456 Street", "city": "Los Angeles"}
	if hasSameAddress(res3, res4) {
		t.Errorf("hasSameAddress passed when it should have failed")
	}
}

func TestUpdateResidenceData(t *testing.T) {
	existing := map[string]interface{}{"residenceId": "123", "streetAddress": "123 Street", "city": "New York"}
	new := map[string]interface{}{"residenceId": "123", "streetAddress": "456 Street", "city": "Los Angeles"}

	updateResidenceData(existing, new)

	if !reflect.DeepEqual(existing, new) {
		t.Errorf("updateResidenceData failed to update the residence data")
	}
}

func TestUpdateResidenceDataAll(t *testing.T) {
	existingResidences := []map[string]interface{}{
		{"residenceId": "123", "streetAddress": "123 Street", "city": "New York", "currentResidence": true},
		{"residenceId": "456", "streetAddress": "456 Street", "city": "Los Angeles", "currentResidence": false},
	}
	newResidence := map[string]interface{}{"residenceId": "789", "streetAddress": "789 Street", "city": "Chicago", "currentResidence": true}

	updatedResidences := UpdateResidenceData(existingResidences, newResidence)

	if len(updatedResidences) != 3 {
		t.Errorf("UpdateResidenceData failed to append the new residence")
	}

	if updatedResidences[0]["currentResidence"].(bool) {
		t.Errorf("UpdateResidenceData failed to set currentResidence of existing residences to false")
	}

	if !updatedResidences[2]["currentResidence"].(bool) {
		t.Errorf("UpdateResidenceData failed to set currentResidence of new residence to true")
	}
}
