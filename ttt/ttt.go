package ttt

//go:generate mockgen -destination=./mocks/clients_repo.go -package=repo_mocks checksmt/ttt IRepo

type IRepo interface {
	Get(id string) (string, error)
	GetTime() (string, error)
}
