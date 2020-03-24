package gfoo

type Iter func(thread *Thread, pos Pos) (Val, error)

func (self Iter) For(action func(val Val, thread *Thread, pos Pos) error, thread *Thread, pos Pos) error {
	for {
		v, err := self(thread, pos)
		
		if err != nil {
			return err
		}
		
		if v == Nil {
			break
		}
		
		if err = action(v, thread, pos); err == &Break {
			break
		} else if err != nil {
			return err
		}
	}

	return nil
}
