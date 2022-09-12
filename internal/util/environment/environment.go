package environment

import "github.com/hizzuu/beatic-backend/conf"

func IsProd() bool {
	return conf.C.App.Environment == "prod"
}

func IsDev() bool {
	return conf.C.App.Environment == "dev"
}

func IsTest() bool {
	return conf.C.App.Environment == "Test"
}
