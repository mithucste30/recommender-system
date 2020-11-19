package main

type RecommenderService interface {
	New(repo Repository) RecommenderService
	Rate(user string, item string, score float64) error
	GetRecommendedItems(user string, count int) ([]string, error)
	//SuggestedItems(user string, max int)([]string, error)
	//SimilarItems(item string, max int)([]string, error)
	//TopItems(user string, max int)([]string, error)
	//AddNewUser(user string) error
}

type recommenderService struct {
	repo Repository
}

func (recommenderService) New(repo Repository) RecommenderService {
	return recommenderService{
		repo: repo,
	}
}

func (svc recommenderService) Rate(user string, item string, score float64) error  {
	return svc.repo.Recommender().Rate(user, item, score)
}

func (svc recommenderService) GetRecommendedItems(user string, count int)([]string, error)  {
	response, err := svc.repo.Recommender().GetUserSuggestions(user, count)
	if err != nil {
		return nil, err
	}
	var  items []string
	for i := 0; i < len(response); i += 2 {
		items = append(items, response[i])
	}
	return items, nil
}