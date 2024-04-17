package service

var manager *transferManager

func init() {
	manager = &transferManager{
		transfers: make(map[string]*transfer),
	}
}