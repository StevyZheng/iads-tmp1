package common

func CheckErr(err error) {
	if err != nil {
		println(err)
	}
}
