func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	email string,
	password string,
	companyID uuid.UUID,
	roleID uuid.UUID,
) error {

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user, err := userdomain.NewUser(
		name,
		email,
		string(hashed),
		companyID,
		roleID,
	)
	if err != nil {
		return err
	}

	return s.userRepo.Create(ctx, user)
}
