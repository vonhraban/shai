package openai_client

import "time"

type DummyClient struct{}

func (c *DummyClient) SetDebug(bool) {}

func (c *DummyClient) PromptCompletions(message string) (string, error) {
	// imitate actual loading time
	time.Sleep(2 * time.Second)

	return "echo 'This is a dummy response; disable DUMMY_API to connect to the real API'", nil
}
