package service


var (
	manager *transferManager
	//directory string = DEFAULT_DIRECTORY
)

func init() {
	manager = &transferManager{
		transfers: make(map[string]*transfer),
	}
}

// func SetOutputDirectory(out_directory string) {
// 	directory = out_directory
// }

// func GetOutputDirectory() string {
// 	return directory
// }