package objects

func builtinIsString(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*String); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsInt(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Int); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsFloat(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Float); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsBool(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Bool); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsChar(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Char); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsBytes(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Bytes); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Array); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsImmutableArray(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*ImmutableArray); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Map); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsImmutableMap(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*ImmutableMap); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsTime(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Time); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsError(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if _, ok := args[0].(*Error); ok {
		return TrueValue, nil
	}

	return FalseValue, nil
}

func builtinIsUndefined(args ...Object) (Object, error) {
	if len(args) != 1 {
		return nil, ErrWrongNumArguments
	}

	if args[0] == UndefinedValue {
		return TrueValue, nil
	}

	return FalseValue, nil
}
