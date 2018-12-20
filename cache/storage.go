package cache

import (
	"io/ioutil"

	"github.com/rumblefrog/ChronoGG-Discord-Bot/sale"
)

func CompareSale(sale *sale.Sale) (bool, error) {
	bytes, err := ioutil.ReadFile(CacheFile)

	if err != nil {
		return false, err
	}

	return string(bytes) == sale.Name, nil
}

func WriteSale(sale *sale.Sale) error {
	if err := ioutil.WriteFile(CacheFile, []byte(sale.Name), 0666); err != nil {
		return err
	}

	return nil
}
