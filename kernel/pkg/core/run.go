package core

func Run(filename string) error {
	if err := loadBin(filename); err != nil {
		return err
	}
	return nil
}
