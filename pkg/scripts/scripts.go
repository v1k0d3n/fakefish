// package scripts

// import (
// 	"log"
// 	"os/exec"
// )

// func ExecuteScript(scriptPath string) error {
// 	cmd := exec.Command("/bin/sh", scriptPath)
// 	output, err := cmd.CombinedOutput()
// 	if err != nil {
// 		log.Printf("Error executing script: %v", err)
// 		return err
// 	}

// 	log.Printf("Script output: %s", output)
// 	return nil
// }

package scripts

import (
	"log"
	"os/exec"
)

func ExecuteScript(scriptPath string) error {
	log.Printf("Executing script: %s", scriptPath)
	cmd := exec.Command("/bin/sh", scriptPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Error executing script: %v", err)
		log.Printf("Script output: %s", output)
		return err
	}

	log.Printf("Script output: %s", output)
	return nil
}
