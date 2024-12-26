t.Run("Testing negative number input", func(t *testing.T) {
    panicFlag := false
	defer func() {
		if r := recover(); r != nil {
			panicFlag = true
		}
	}()

	RandomEmail(-10)

    if !panicFlag {
        t.Errorf("Code should have panicked")
    } else {
	    t.Logf("Handled negative input gracefully with a panic")
    }
})
