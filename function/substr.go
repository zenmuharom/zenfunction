package function

func (assigner *DefaultAssigner) Substr(arg string, from int, to int) (result string, err error) {
	to = from + to

	if to == 0 || to > len(arg) {
		to = len(arg)
	}

	result = arg[from:to]

	return
}
