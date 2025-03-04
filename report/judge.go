package report

type Judge struct {
	*Report
}

func (p Judge) Beauty() {
	p.Info("[Comparison Result] ")
	switch p.d.IsAccept() {
	case true:
		p.Successln("✅ Accept (Output Matched)")
	case false:
		p.Errorln("❌ Error (Output Mismatch)")
	}
}
