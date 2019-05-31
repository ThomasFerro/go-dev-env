package triggers

type Trigger interface {
	Init() chan bool
}
