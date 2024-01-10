package openai_client

type Client interface {
	SetDebug(bool)
	PromptCompletions(message string) (string, error)
}
