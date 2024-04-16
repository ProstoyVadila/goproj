package clipboard

import (
	cb "golang.design/x/clipboard"
)

func Save(cmd string) error {
	err := cb.Init()
	if err != nil {
		return err
	}

	cb.Write(cb.FmtText, []byte(cmd))
	return nil
}
