package loader

type Service interface {
	ParseDocument(url string) (successCount, failedCount int64, err error)
}
