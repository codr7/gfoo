package gfoo

type TimeDelta struct {
     years, months, days int
}

func Days(n int) TimeDelta {
     return TimeDelta{days: n}
}

func (self TimeDelta) Compare(other TimeDelta) Order {
	if out := CompareInt(self.years, other.years); out != Eq {
		return out
	}

	if out := CompareInt(self.months, other.months); out != Eq {
		return out
	}

	return CompareInt(self.days, other.days)
}
