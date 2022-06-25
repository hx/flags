package states

type Machine interface {
	Get() State
	Set(state State)
	Toggle(index int)
}
