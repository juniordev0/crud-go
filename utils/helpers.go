package utils

import (
	"database/sql/driver"
	"fmt"
	"strings"
	"time"
)

// Tipo personalizado para formatar a data
type Date struct {
	time.Time
}

// Implementar o método Scan
func (d *Date) Scan(value interface{}) error {
	if value == nil {
		*d = Date{}
		return nil
	}
	t, ok := value.(time.Time)
	if !ok {
		return fmt.Errorf("falha ao converter %v para time.Time", value)
	}
	d.Time = t
	return nil
}

// Implementar o método Value (opcional, para inserts)
func (d Date) Value() (driver.Value, error) {
	return d.Time, nil
}

// Formata a data como string no formato "2006-01-02"
func (d Date) MarshalJSON() ([]byte, error) {
	formatted := d.Format("2006-01-02")
	return []byte(`"` + formatted + `"`), nil
}

// Permite interpretar datas em diferentes formatos
func (d *Date) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), "\"")
	if str == "" || str == "null" {
		*d = Date{}
		return nil
	}

	layouts := []string{"2006-01-02", "01-02-2006", "01/02/2006"}
	var err error
	for _, layout := range layouts {
		var t time.Time
		t, err = time.Parse(layout, str)
		if err == nil {
			d.Time = t
			return nil
		}
	}
	return fmt.Errorf("formato de data inválido: %s", str)
}
