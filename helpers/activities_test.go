package helpers

import (
	"reflect"
	"testing"
)

func TestUpdateActivityData(t *testing.T) {
	tests := []struct {
		name          string
		existing      map[string]map[string]interface{}
		newActivity   map[string]interface{}
		want          map[string]map[string]interface{}
		wantErr       bool
		expectedError string
	}{
		{
			name:     "Successful Update",
			existing: map[string]map[string]interface{}{"oldActivity": {"oldKey": "oldValue"}},
			newActivity: map[string]interface{}{
				"newActivity": map[string]interface{}{
					"newKey": "newValue",
				},
			},
			want:    map[string]map[string]interface{}{"oldActivity": {"oldKey": "oldValue"}, "newActivity": {"newKey": "newValue"}},
			wantErr: false,
		},
		{
			name:     "Error: Empty Key",
			existing: map[string]map[string]interface{}{},
			newActivity: map[string]interface{}{
				"": map[string]interface{}{
					"newKey": "newValue",
				},
			},
			wantErr:       true,
			expectedError: "Key is empty",
		},
		{
			name:     "Error: Incorrect JSON Format",
			existing: map[string]map[string]interface{}{},
			newActivity: map[string]interface{}{
				"newActivity": "not a map",
			},
			wantErr:       true,
			expectedError: "JSON for activity is not formatted correctly",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := UpdateActivityData(tt.existing, tt.newActivity)

			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateActivityData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr && err.Error() != tt.expectedError {
				t.Errorf("UpdateActivityData() error message = %v, expectedError %v", err.Error(), tt.expectedError)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateActivityData() = %v, want %v", got, tt.want)
			}
		})
	}
}
