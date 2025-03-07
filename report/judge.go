package report

type Judge struct {
	*Report
}

func (p Judge) Beauty() {
	p.Info("[Comparison Result] ")
	switch p.d.IsAccept() {
	case true:
		p.Successln("✅ Accepted! Your output matches the expected result.")
	case false:
		p.Errorln("❌ Incorrect output! The result does not match the expected output.")
	}
}
