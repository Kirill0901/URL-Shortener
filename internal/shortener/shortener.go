package shortener

var count int64 = 0

var alph = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

type CountGetter interface {
	GetCount() (int64, error)
}

func CountInit(countGetter CountGetter) error {

	num, err := countGetter.GetCount()
	if err != nil {
		count = 0
		return err
	}

	count = num

	return nil
}

func MakeShorter(long_url string) (string, error) {
	short_url := ""

	n := count

	for i := 0; i < 10; i++ {
		short_url += string(alph[n%63])
		n /= 63
	}

	count += 1

	return short_url, nil
}
