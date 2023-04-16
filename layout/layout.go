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

type candidate struct {
	id        int64
	top       [][]int64
	bottom    [][]int64
	isRotated bool
}

type Direction int

const (
	Top Direction = iota
	Bottom
)

func Layout(blocks []Block) []LayoutResult {
	if len(blocks) == 0 {
		return make([]LayoutResult, 0)
	}

	cand := make([]candidate, 0)

	// put original blocks
	for _, b := range blocks {
		cand = append(cand, candidate{
			id:        b.Id,
			top:       openBlock(b.Form, Top),
			bottom:    openBlock(b.Form, Bottom),
			isRotated: false,
		})
	}

	// put rotated blocks
	for i := 0; i < len(blocks); i++ {
		cand = append(cand, candidate{
			id:        cand[i].id,
			top:       rotate(cand[i].bottom),
			bottom:    rotate(cand[i].top),
			isRotated: true,
		})
	}

	path := dfs(len(blocks), cand, make(map[int64]bool), make([]LayoutResult, 0), make([][]int64, 0))
	if path == nil {
		panic("unexpected empty result")
	}

	return *path
}

func dfs(remaining int, blocks []candidate, used map[int64]bool, path []LayoutResult, top [][]int64) *[]LayoutResult {
	if remaining == 0 {
		return &path
	}

	for _, block := range blocks {
		if _, ok := used[block.id]; !ok {
			if !fit(top, block.bottom) {
				continue
			}

			used[block.id] = true
			solved := dfs(remaining-1, blocks, used, append(path, LayoutResult{
				BlockId:   block.id,
				Position:  len(path) + 1,
				IsRotated: block.isRotated,
			}), block.top)
			if solved != nil {
				return solved
			}
			delete(used, block.id)
		}
	}

	return nil
}

func fit(low [][]int64, high [][]int64) bool {
	if len(low) != len(high) {
		return false
	}

	for i := 0; i < len(low); i++ {
		for j := 0; j < len(low[i]); j++ {
			if low[i][j] == high[i][j] {
				return false
			}
		}
	}

	return true
}

func openBlock(form [][]int64, dir Direction) [][]int64 {
	switch dir {
	case Top:
		var end int
		for end = 0; end < len(form); end++ {
			if !zeros(form[end]) {
				break
			}
		}
		return form[0:end]
	case Bottom:
		var start int
		for start = len(form) - 1; start >= 0; start-- {
			if !zeros(form[start]) {
				break
			}
		}
		return form[start+1:]
	default:
		panic("unknown position arg")
	}
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
	if l == 0 {
		return make([][]int64, 0)
	}

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
