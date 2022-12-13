package account

type CategoriesAccount string

const (
	Master    CategoriesAccount = "master"
	Admin     CategoriesAccount = "admin"
	Producer  CategoriesAccount = "producer"
	Consumer  CategoriesAccount = "consumer"
	Miners    CategoriesAccount = "miniers"
	Evaluator CategoriesAccount = "evaluator"
)
