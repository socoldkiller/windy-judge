package render

type DefaultSFCaseHook struct {
	info SFCaseInfo
}

func (d *DefaultSFCaseHook) SuccessCaseHook(result TestCaseResult) {
	d.info.Passed += 1
	d.info.UsedTime += result.Elapsed
	d.info.SuccessResult = append(d.info.SuccessResult, result)
}

func (d *DefaultSFCaseHook) FailCaseHook(result TestCaseResult) {
	d.info.Failed += 1
	d.info.UsedTime += result.Elapsed
	d.info.FailedResult = append(d.info.FailedResult, result)
	exitCode = 1
}

func (d *DefaultSFCaseHook) Info() SFCaseInfo {
	d.info.Total = d.info.Failed + d.info.Passed
	return d.info
}
