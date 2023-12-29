package entity

import "strconv"

type FizzBuzz struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func NewFizzBuzzEntity(int1, int2, limit int, str1, str2 string) (*FizzBuzz, error) {
	fizzbuzz := &FizzBuzz{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	if err := fizzbuzz.validate(); err != nil {
		return nil, err
	}

	return fizzbuzz, nil
}

func (fb *FizzBuzz) validate() error {
	if fb.Str1 == "" || fb.Str2 == "" {
		return ErrInvalidStringInput
	}

	if fb.Int1 <= 0 || fb.Int2 <= 0 || fb.Limit <= 0 {
		return ErrInvalidIntegerInput
	}

	return nil
}

func (fb *FizzBuzz) CalculateFizzBuzz() []string {
	var ret []string

	for i := 1; i <= fb.Limit; i++ {
		tmp := ""

		if i%fb.Int1 == 0 {
			tmp += fb.Str1
		}

		if i%fb.Int2 == 0 {
			tmp += fb.Str2
		}

		if tmp == "" {
			tmp = strconv.Itoa(i)
		}

		ret = append(ret, tmp)
	}

	return ret
}
