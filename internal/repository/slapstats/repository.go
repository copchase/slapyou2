package slapstats

type slapStatsRepo struct {
	dynamoClient dynamoClient
}

type dynamoClient interface {
}

func NewSlapStatsRepo(dynamoClient dynamoClient) *slapStatsRepo {
	return &slapStatsRepo{
		dynamoClient: dynamoClient,
	}
}

func (s *slapStatsRepo) UpdateUser() error {
	return nil
}
