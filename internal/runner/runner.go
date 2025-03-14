package runner

type Runner[I, O any] interface {
	Run(input I) O
}

type ContextAwareRunner[I, O any] interface {
	PreRun(input I)
	PostRun(input I, output O)
}

type ContextualRunner[I, O any] struct {
	runner    Runner[I, O]
	ctxRunner ContextAwareRunner[I, O]
}

func (r ContextualRunner[I, O]) Run(input I) O {
	var (
		ctxRunner = r.ctxRunner
		runner    = r.runner
	)

	ctxRunner.PreRun(input)
	output := runner.Run(input)
	ctxRunner.PostRun(input, output)
	return output
}

func NewContextualRunner[I, O any](runner Runner[I, O], ctxRunner ContextAwareRunner[I, O]) Runner[I, O] {
	return &ContextualRunner[I, O]{
		runner:    runner,
		ctxRunner: ctxRunner,
	}
}

// BatchContextualRunner is responsible for processing a batch of inputs.
// It leverages a Runner to execute individual tasks and a ContextAwareRunner
// to handle pre-run and post-run operations on the entire batch.
// Their order is as follows: PreRun -> Run -> PostRun.
type BatchContextualRunner[I, O any] struct {
	// runner is the core Runner that processes each individual task.
	runner Runner[I, O]

	// ctxRunner is the ContextAwareRunner that performs pre-run and post-run
	// actions on a slice of inputs and outputs.
	ctxRunner ContextAwareRunner[[]I, []O]
}

func NewBatchContextualRunner[I, O any](runner Runner[I, O], ctxRunner ContextAwareRunner[[]I, []O]) Runner[[]I, []O] {
	rs := &BatchContextualRunner[I, O]{
		runner:    runner,
		ctxRunner: ctxRunner,
	}
	return rs
}

func (r BatchContextualRunner[I, O]) Run(inputs []I) []O {
	var (
		ctxRunner = r.ctxRunner
		runner    = r.runner
		outputs   []O
	)

	ctxRunner.PreRun(inputs)
	for _, in := range inputs {
		output := runner.Run(in)
		outputs = append(outputs, output)
	}

	ctxRunner.PostRun(inputs, outputs)
	return outputs
}
