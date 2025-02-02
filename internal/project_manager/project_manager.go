package project_manager

func UserPromptLoop(prompt string) {
	// TODO: Should continuously listen for user input until the loop is explicitly broken by the user,
	// 	at which point start the next step of code generation loop
	//reader := bufio.NewReader(os.Stdin)
	//for {
	//	input, _ := reader.ReadString('\n')
	//}
}

func DoRequestToRefiner(prompt string) {
	// TODO: Send the request to the Refiner Agent, listen for response.
	//  Pass response to semantic parser to determine if more action is needed.
	//  If not, return response to User, otherwise Semantic Parser will return requests to be made. Make
	//  the requests and send results back to the Refiner Agent. Loop continues until no more semantic parsing
	//  is needed and control defers back to the UserPromptLoop.
}

func DoRequestSemanticParser(prompt string) []byte {
	// TODO:
	return nil
}
