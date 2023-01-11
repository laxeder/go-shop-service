package str

func MixStrings(str1, str2 string) (mix string) {

	mix = ""

	lenStr1 := len(str1)
	lenStr2 := len(str2)

	lenMix := (lenStr1 + lenStr2)

	for i := 1; i <= lenMix; i++ {

		if i <= lenStr1 {
			mix += str1[(i - 1):i]
		}

		if i <= lenStr2 {
			mix += str2[(i - 1):i]
		}

	}

	return
}
