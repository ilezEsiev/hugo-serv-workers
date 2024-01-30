package counter

import (
	"fmt"
	"log"
	"os"
	"time"
)

type FileWriter interface {
	WriteFile(filename string, data []byte, perm os.FileMode) error
}

type RealFileWriter struct{}

func (r *RealFileWriter) WriteFile(filename string, data []byte, perm os.FileMode) error {
	return os.WriteFile(filename, data, perm)
}

var fileWriter FileWriter = &RealFileWriter{}

const content = `---
menu:
    before:
        name: tasks
        weight: 5
title: Обновление данных в реальном времени
---

# Задача: Обновление данных в реальном времени

Напишите воркер, который будет обновлять данные в реальном времени, на текущей странице.
Текст данной задачи менять нельзя, только время и счетчик.

Файл данной страницы: %v/app/static/tasks/_index.md%v

Должен меняться счетчик и время:

Текущее время: %v

Счетчик: %v



## Критерии приемки:
- [ ] Воркер должен обновлять данные каждые 5 секунд
- [ ] Счетчик должен увеличиваться на 1 каждые 5 секунд
- [ ] Время должно обновляться каждые 5 секунд`

const (
	appPathToCounterFile = `/app/static/tasks/_index.md`
)

func CounterWorker() {
	t := time.NewTicker(5 * time.Second)
	var b byte = 0
	for {
		select {
		case <-t.C:
			err := fileWriter.WriteFile(appPathToCounterFile, []byte(fmt.Sprintf(content, "`", "`", time.Now().Format("2006-01-02 15:04:05"), b)), 0644)
			if err != nil {
				log.Println(err)
			}
			b++
		}
	}
}
