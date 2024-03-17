package change_actor_data

import (
	"time"
)

func validateInParams(in *ChangeActorDataIn) (bool, string) {
	if in.ID < 0 {
		return false, "actor id should be positive number"
	}

	if in.Gender != "" && in.Gender != "male" && in.Gender != "female" {
		return false, "invalid gender"
	}

	if in.BirthDate != "" {
		t, err := time.Parse("2006-01-02", in.BirthDate)
		if err != nil {
			return false, "invalid birth date format"
		}

		if !t.IsZero() && t.After(time.Now()) {
			return false, "birth date should be before current time"
		}
	}

	return true, ""
}
