package ports

type ExtractWallet interface {
	GenerateExtract(id string) error
}
