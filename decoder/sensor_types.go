package decoder





func IsValidCategory(category string) (bool, error) {
	d := SensorNames
	switch category {
	case
		string(d.ME),
		string(d.ME),
		string(d.ME),
		string(d.ME):
		return true, nil
	}
	return false, nil
}

