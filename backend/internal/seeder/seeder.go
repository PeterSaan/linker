package seeder

func Main() error {
	if err := UserSeeder(); err != nil {
		return err
	}

	if err := ProfileSeeder(); err != nil {
		return err
	}

	return nil
}
