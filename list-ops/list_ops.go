package listops

type binFunc func(int, int) int
type predFunc func(int) bool
type unaryFunc func(int) int

type IntList []int

func (l IntList) Foldl(fn binFunc, initial int) int {
	for _, v := range l {
		initial = fn(initial, v)
	}
	return initial
}

func (l IntList) Foldr(fn binFunc, initial int) int {
	for i := l.Length() - 1; i >= 0; i-- {
		initial = fn(l[i], initial)
	}
	return initial
}

func (l IntList) Filter(fn predFunc) IntList {
	for i, v := range l {
		if !fn(v) {
			l = IntList(l[:i]).Append(IntList(l[i+1:]))
		}
	}
	return l
}

func (l IntList) Length() int {
	return len(l)
}

func (l IntList) Map(fn unaryFunc) IntList {
	for i, v := range l {
		l[i] = fn(v)
	}
	return l
}

func (l IntList) Reverse() IntList {
	for i, j := 0, l.Length()-1; i < j; i, j = i+1, j-1 {
		l[i], l[j] = l[j], l[i]
	}
	return l
}

func (l1 IntList) Append(l2 IntList) IntList {
	return append(l1, l2...)
}

func (l1 IntList) Concat(ls []IntList) IntList {
	var l2 IntList
	for _, l := range ls {
		l2 = l2.Append(l)
	}
	return l1.Append(l2)
}
