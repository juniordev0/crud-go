package utils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Tipo personalizado para formatar a data
type Date struct {
	time.Time
}

// Implementar o método Scan
func (d *Date) Scan(value interface{}) error {
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
