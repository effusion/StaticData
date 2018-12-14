package xlsximporter

type ImportResult struct{}

//
type Importer interface {
	ImportStaticDataXlsx(filePath string) (result *ImportResult)
}
