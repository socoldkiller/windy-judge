package renderTask

import (
	"fmt"
	"strconv"
	"windy-judge/internal/F"
)

type DefaultTestResult struct {
	*TestCaseResultCount
	p F.Printer
}

func (d *DefaultTestResult) Write(data any) error {
	res := data.([]TestCaseResult)
	d.total = len(res)
	return nil
}

func (d *DefaultTestResult) Beauty() {
	p := d.p
	d.failed = d.total - d.passed
	usedTime := fmt.Sprintf("%.2fs", d.usedTime.Seconds())
	switch {
	case d.total == 0:
		p.Defaultln("🔍 No test cases were found or executed. Please ensure your tests are correctly set up. 🛠️")
	case d.failed == 0:
		p.Defaultln("🎉 Congratulations! All "+strconv.Itoa(d.total)+" test cases passed successfully! ✅🎯", "Execution time:", usedTime, "Keep up the great work! 🚀🔥")
	default:
		p.Errorln("❌ "+strconv.Itoa(d.failed)+"/"+strconv.Itoa(d.total)+" test cases failed!", "Execution time:", usedTime, "Please check the errors above.")
	}
}
