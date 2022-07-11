package pkg

func (a Uint) Compare(b Uint) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Uint8) Compare(b Uint8) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Uint16) Compare(b Uint16) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Uint32) Compare(b Uint32) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Uint64) Compare(b Uint64) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Int) Compare(b Int) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Int8) Compare(b Int8) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Int16) Compare(b Int16) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Int32) Compare(b Int32) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Int64) Compare(b Int64) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Float32) Compare(b Float32) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a Float64) Compare(b Float64) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}

func (a String) Compare(b String) Ordering {
	if a < b {
		return OrderingLessThan
	} else if a == b {
		return OrderingEqual
	} else {
		return OrderingGreaterThan
	}
}
