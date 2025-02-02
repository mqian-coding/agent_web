package project_manager

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

var (
	Host = "localhost"
)

const (
	DeepseekCoderModelName = "deepseek-coder:33b"
)

type OllamaResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

func UserPromptLoop(prompt string) {
	// TODO: Should continuously listen for user input until the loop is explicitly broken by the user,
	// 	at which point start the next step of code generation loop
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		DoRequestToRefiner(input)
	}
}

func DoRequestToRefiner(prompt string) {
	// TODO: Send the request to the Refiner Agent, listen for response.
	//  Pass response to semantic parser to determine if more action is needed.
	//  If not, return response to User, otherwise Semantic Parser will return requests to be made. Make
	//  the requests and send results back to the Refiner Agent. Loop continues until no more semantic parsing
	//  is needed and control defers back to the UserPromptLoop.

}

func DoRequestSemanticParser(prompt string) []byte {
	// TODO: Forward the response from the refiner,
	return nil
}

func DoRequestToAgent(prompt, model, port string) (string, error) {
	var err error
	body := bytes.NewReader([]byte(fmt.Sprintf("{\n\t\"model\": \"%s\",\n\t\"prompt\":\"%s\"\n}", model, prompt)))
	url := fmt.Sprintf("http://%s:%s/api/generate", Host, port)
	req, err := http.NewRequest(http.MethodPost, url, body)
	if err != nil {
		return "", err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	nlResp, err := AgentResponseToNaturalLanguage(resp)
	return nlResp, err
}

func AgentResponseToNaturalLanguage(resp *http.Response) (string, error) {
	if resp == nil {
		return "", errors.New("resp *http.Response is nil")
	}
	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	responses := strings.Split(string(raw), "\n")
	builder := &strings.Builder{}
	ollamaResp := &OllamaResponse{}
	for _, response := range responses {
		json.Unmarshal([]byte(response), ollamaResp)
		builder.Write([]byte(ollamaResp.Response))
	}
	return builder.String(), nil
}
