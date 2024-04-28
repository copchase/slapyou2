package slap

type slapService struct {
	slapStatsRepo slapStatsRepo
}

type slapStatsRepo interface {
	UpdateUser() error
}

func NewSlapService(slapStatsRepo slapStatsRepo) *slapService {
	return &slapService{
		slapStatsRepo: slapStatsRepo,
	}
}
