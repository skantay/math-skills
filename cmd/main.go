package main

import (
	"fmt"

	"github.com/skantay/math-skills/internal/mathSkills"
)

func main() {
	if err := mathSkills.Run(); err != nil {
		//nolint:forbidigo
		fmt.Printf("%s\n\n%s\n", mathSkills.Error(), err)
	}
}
