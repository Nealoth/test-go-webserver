package configuration

type Reader interface {
	read() (*GeneralConfiguration, error)
}

func ReadConfiguration(reader Reader) (*GeneralConfiguration, error) {
	return reader.read()
}
