package main

/**
 * We need to get a reference to the
 * object we are adapting.
 */
type turkeyAdapter struct {
	turkey turkey
}

func newTurkeyAdaptor(t turkey) *turkeyAdapter {
	return &turkeyAdapter{
		turkey: t,
	}
}

/**
 * Implement the interface of the
 * type you’re adapting to. This is the
 * interface your client expects to see
 */
func (t *turkeyAdapter) quack() {
	t.turkey.gobble()
}

func (t *turkeyAdapter) fly() {
	for i := 1; i <= 5; i++ {
		t.turkey.fly()
	}
}
