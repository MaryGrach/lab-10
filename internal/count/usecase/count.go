package usecase

func (u *Usecase) FetchCountMessage() (int, error) {
	msg, err := u.p.SelectCount()
	if err != nil {
		return 0, err
	}

	return msg, nil
}

func (u *Usecase) UpdateCountMessage(msg int) error {

	err := u.p.UpdateCount(msg)
	if err != nil {
		return err
	}

	return nil
}
