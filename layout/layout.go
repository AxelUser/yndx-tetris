package layout

type Block struct {
	Id   int64
	Form [][]int64
}

type LayoutResult struct {
	BlockId   int64
	Position  int
	IsRotated bool
}

func Layout(blocks []Block) []LayoutResult {
	if len(blocks) == 0 {
		return make([]LayoutResult, 0)
	}

	cand := make([]candidate, 0)
	for _, b := range blocks {
		cand = append(cand, candidate{
			id:        b.Id,
			form:      b.Form,
			isRotated: false,
		})
	}

	for _, b := range blocks {
		cand = append(cand, candidate{
			id:        b.Id,
			form:      rotate(b.Form),
			isRotated: true,
		})
	}

	path := dfs(len(blocks), cand, make(map[int64]bool), make([]candidate, 0), make([][]int64, 0))
	if path == nil {
		panic("unexpected empty result")
	}

	res := make([]LayoutResult, 0)
	for i, c := range *path {
		res = append(res, LayoutResult{
			BlockId:   c.id,
			Position:  i + 1,
			IsRotated: c.isRotated,
		})
	}

	return res
}

type candidate struct {
	id        int64
	form      [][]int64
	isRotated bool
}

func dfs(remaining int, blocks []candidate, used map[int64]bool, path []candidate, merged [][]int64) *[]candidate {
	if remaining == 0 {
		return &path
	}

	for _, block := range blocks {
		if _, ok := used[block.id]; !ok {
			m := merge(merged, block.form)
			if m == nil {
				continue
			}

			used[block.id] = true
			solved := dfs(remaining-1, blocks, used, append(path, block), *m)
			if solved != nil {
				return solved
			}
			delete(used, block.id)
		}
	}

	return nil
}

func merge(low [][]int64, high [][]int64) *[][]int64 {
	zlow := openBlock(low, true)
	zhigh := openBlock(high, false)
	if len(zlow) != len(zhigh) {
		return nil
	}

	for i := 0; i < len(zlow); i++ {
		for j := 0; j < len(zlow[i]); j++ {
			if zlow[i][j] == zhigh[i][j] {
				return nil
			}
		}
	}

	merged := high[0 : len(high)-len(zhigh)]

	return &merged
}

func openBlock(form [][]int64, topToLow bool) [][]int64 {
	open := make([][]int64, 0)

	if topToLow {
		for _, line := range form {
			if !zeros(line) {
				break
			}

			open = append(open, line)
		}
	} else {
		for i := len(form) - 1; i >= 0; i-- {
			if !zeros(form[i]) {
				break
			}

			open = append([][]int64{form[i]}, open...)
		}
	}

	return open
}

func zeros(line []int64) bool {
	for _, v := range line {
		if v == 0 {
			return true
		}
	}

	return false
}

func rotate(form [][]int64) [][]int64 {
	l := len(form)
	w := len(form[0])
	r := make([][]int64, l)
	for i := range form {
		rline := make([]int64, w)
		for j, v := range form[i] {
			rline[w-j-1] = v
		}

		r[l-i-1] = rline
	}

	return r
}
