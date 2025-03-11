package report

import "time"

type Title struct {
	*Report
}

func (p Title) Beauty() {
	p.Infoln("# Test Case", p.ID, "- Result:")
	p.Info("----------------------------------------------\n")
	p.Timeln("[Timestamps]")
	p.KeyValueFormat("%s\n", "- Test Time: ", time.Now().Format(time.DateTime))
	p.KeyValueFormat("%.2fs\n", "- Execution Time: ", p.ExecutionTime.Seconds())
}
