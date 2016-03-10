package job

var (
	NewID = newID()
)

func newID() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}
