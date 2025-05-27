package script

import "os"

// ReadFile implements https://book.getfoundry.sh/cheatcodes/fs
func (c *CheatCodesPrecompile) ReadFile(input string) ([]byte, error) {
	data, err := os.ReadFile(input)
	if err != nil {
		return nil, err
	}
	return data, nil
}
