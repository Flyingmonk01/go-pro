package helper

func CheckNilError(err error) (error, bool) {
	if err != nil {
		return err, true
	}
	return nil, false
}
