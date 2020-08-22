package sensitive

var (
	acAutoMachine *AcAutoMachine
)

func AllowPublish(content string) bool {
	return !acAutoMachine.Match(content)
}
