package helpers

func getKeyValueFromJSON(jsonData map[string]interface{}) (string, any) {
    for key, value := range jsonData {
        return key, value
    }

    return "", nil
}
