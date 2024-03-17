package create_actor

import (
	"fmt"
	"time"
)

func validateInParams(in *CreateActorIn) (bool, string) {
	if len(in.Name) == 0 || len(in.Gender) == 0 {
		return false, fmt.Sprintf("the name or gender should not be empty")
	}

	if in.Gender != "male" && in.Gender != "female" {
		return false, "invalid gender"
	}

	t, err := time.Parse("2006-01-02", in.BirthDate)
	if err != nil {
		return false, "invalid birth date format"
	}

	if !t.IsZero() && t.After(time.Now()) {
		return false, "birth date should be before current time"
	}

	return true, ""
}
