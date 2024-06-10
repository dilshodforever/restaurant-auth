package config

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
