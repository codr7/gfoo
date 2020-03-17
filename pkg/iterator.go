package gfoo

type Iterator func(scope *Scope, pos Pos) (*Val, error)

func (self Iterator) For(action func(val Val, scope *Scope, pos Pos) error, scope *Scope, pos Pos) error {
	for {			
		v, err := self(scope, pos)
		
		if err != nil {
			return err
		}
		
		if v == nil {
			break
		}
		
		if *v == Nil {
			continue
		}
		
		if err = action(*v, scope, pos); err != nil {
			return err
		}
	}

	return nil
}
