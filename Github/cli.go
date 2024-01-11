package Github

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type CLI struct {
	Repo string
}

func GetPulls(cli *CLI) []Pull {
	endpoint := fmt.Sprintf("%s/pulls", cli.Repo)
	args := []string{"api", "--method", "GET", endpoint}
	cmd := exec.Command("gh", args...)

	response, err := cmd.CombinedOutput()
	//fmt.Println(string(response))
	if err != nil {
		fmt.Println("Error:", err)
		return []Pull{}
	}

	var pulls []Pull

	err = json.Unmarshal(response, &pulls)
	if err != nil {
		fmt.Println("Error:", err)
		return []Pull{}
	}

	if err != nil {
		fmt.Println("Error:", err)
		return []Pull{}
	}

	return pulls
}
