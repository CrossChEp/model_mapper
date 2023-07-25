package model_mapper

func Map(to interface{}, from interface{}, skipNulls bool) error {
	if !skipNulls {
		if err := mapWithNullFields(to, from); err != nil {

		}
		return nil
	}
	return nil
}
