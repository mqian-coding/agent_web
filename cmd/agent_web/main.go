package main

import (
	"fmt"
	"github.com/mqian-coding/agent_web/internal/project_manager"
)

func main() {
	resp, err := project_manager.DoRequestToAgent("Generate a bash file called hello_world.sh which just echos hello world!", project_manager.DeepseekCoderModelName, "11434")
	fmt.Println(resp)
	if err != nil {
		fmt.Println(err.Error())
	}
}
