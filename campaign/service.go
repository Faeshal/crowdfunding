package campaign

type Service interface {
	FindCampaigns(userID int) ([]Campaign, error)
}

// * butuh dependecy ke repository
type service struct {
	repository Repository
}

// * supaya bisa di call dari luar package
func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindCampaigns(userID int) ([]Campaign, error) {

	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}
		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}
