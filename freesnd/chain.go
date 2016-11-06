package main

// chain chains some command funcs together.
func chain(cmdfuncs ...commandFunc) commandFunc {
	return func(args []string) error {
		for _, f := range cmdfuncs {
			if err := f(args); err != nil {
				return err
			}
		}
		return nil
	}
}
