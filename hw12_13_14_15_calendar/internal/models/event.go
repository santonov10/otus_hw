package models

import "time"

type Event struct {
	ID          string    // уникальный идентификатор события (можно воспользоваться UUID);
	Title       string    // Заголовок - короткий текст;
	StartTime   time.Time // Дата и время события;
	EndTime     time.Time // время окончания;
	Description string    // Описание события - длинный текст, опционально;
	UserID      string    // id владельца события;
}
