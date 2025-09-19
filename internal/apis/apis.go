//go:generate go run go.uber.org/mock/mockgen -package=apistest -source=apis.go -destination=apistest/apis.go .

package apis

type Github interface {
	AuthenticateAndWriteComment(fullRepoName string, issueNumber int) error
}
