package user

type RepositoryUser interface {
	Register(*UserParams) Response
	Login(*UserParams) Response
}

func Register(account, pwd, groupId string) {

}
