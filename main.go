package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/compute/v1"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$$ ")
		cmdString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		err = runCommand(cmdString)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

}
func runCommand(commandStr string) error {
	commandStr = strings.TrimSuffix(commandStr, "\n")
	arrCommandStr := strings.Fields(commandStr) // formatting command input
	switch arrCommandStr[0] {
	case "exit":
		os.Exit(0)
	case "plus":
		if len(arrCommandStr) < 3 {
			return errors.New("Required: 2 arguments")
		}
		arrNum := []int64{}
		for i, arg := range arrCommandStr {
			if i == 0 {
				continue
			}
			n, _ := strconv.ParseInt(arg, 10, 64)
			arrNum = append(arrNum, n)
		}
		fmt.Fprintln(os.Stdout, sum(arrNum...))
		return nil
	case "gcloud": //not complete

		ctx := context.Background()
		c, err := google.DefaultClient(ctx, compute.CloudPlatformScope)
		if err != nil {
			log.Fatal(err)
		}
		computeService, err := compute.New(c)
		if err != nil {
			log.Fatal(err)
		}
		project := "powerful-decker-113215"
		zone := "us-central1-b"
		instance := "learning"

		resp, err := computeService.Instances.Start(project, zone, instance).Context(ctx).Do()
		if err != nil {
			log.Fatal(err)
		}
		// change below code to process resp object
		fmt.Printf("%#v\n", resp)

	}
	cmd := exec.Command(arrCommandStr[0], arrCommandStr[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}

// definition of some simple test command
func sum(numbers ...int64) int64 {
	res := int64(0)
	for _, num := range numbers {
		res += num
	}
	return res
}
