package _type

type GoQuery string

func (q GoQuery) ToString() string {
	return string(q)
}
