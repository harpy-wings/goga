package goga

type modelRecord struct {
	model Model
	cost  float64
}

type modelSortedList []modelRecord

func (a modelSortedList) Len() int           { return len(a) }
func (a modelSortedList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a modelSortedList) Less(i, j int) bool { return a[i].cost < a[j].cost }
