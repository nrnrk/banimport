package inner

func IsPrime(n uint) bool {
	for i := uint(0); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
