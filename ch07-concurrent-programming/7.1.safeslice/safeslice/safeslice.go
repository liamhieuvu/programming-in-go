package safeslice

type safeSlice chan commandData

type commandData struct {
	action  commandAction
	index   int
	value   interface{}
	result  chan<- interface{}
	data    chan<- []interface{}
	updater UpdateFunc
}

type commandAction int

const (
	insert commandAction = iota
	at
	end
	remove
	length
	update
)

type SafeSlice interface {
	Append(interface{})
	At(int) interface{}
	Close() []interface{}
	Delete(int)
	Len() int
	Update(int, UpdateFunc)
}

type UpdateFunc func(interface{}) interface{}

func New() SafeSlice {
	ss := make(safeSlice)
	go ss.run()
	return ss
}

func (ss safeSlice) run() {
	store := make([]interface{}, 0)
	for command := range ss {
		switch command.action {
		case insert:
			store = append(store, command.value)
		case at:
			if 0 <= command.index && command.index < len(store) {
				command.result <- store[command.index]
			} else {
				command.result <- nil
			}
		case end:
			close(ss)
			command.data <- store
		case remove:
			if 0 <= command.index && command.index < len(store) {
				store = append(store[:command.index], store[command.index+1:]...)
			}
		case length:
			command.result <- len(store)
		case update:
			if 0 <= command.index && command.index < len(store) {
				store[command.index] = command.updater(store[command.index])
			}
		}
	}
}

func (ss safeSlice) Append(value interface{}) {
	ss <- commandData{action: insert, value: value}
}

func (ss safeSlice) At(index int) interface{} {
	reply := make(chan interface{})
	ss <- commandData{action: at, index: index, result: reply}
	return <-reply
}

func (ss safeSlice) Close() []interface{} {
	reply := make(chan []interface{})
	ss <- commandData{action: end, data: reply}
	return <-reply
}

func (ss safeSlice) Delete(index int) {
	ss <- commandData{action: remove, index: index}
}

func (ss safeSlice) Len() int {
	reply := make(chan interface{})
	ss <- commandData{action: length, result: reply}
	return (<-reply).(int)
}

// If the updater calls a safeSlice method we will get deadlock!
func (ss safeSlice) Update(index int, updater UpdateFunc) {
	ss <- commandData{action: update, index: index, updater: updater}
}
