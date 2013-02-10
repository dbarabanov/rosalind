package util

type StringIterator func() (letter byte, ok bool)

func MakeStringIterator(s string) StringIterator {
	i := -1
	var letters = []byte(s)
	return func() (byte, bool) {
		for i+1 < len(letters) {
			i++
			return s[i], true
		}
		return 0, false
	}
}
