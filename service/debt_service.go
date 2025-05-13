package service

import (
	"sync"

	"github.com/google/uuid"
)

var (
	debts = make(map[string]string)
	mutex sync.Mutex
)

func SaveRawDebt(jsonBody string) string {
	id := uuid.New().String()

	mutex.Lock()
	defer mutex.Unlock()
	debts[id] = jsonBody

	return id
}

func GetDebtById(id string) (string, bool) {
	mutex.Lock()
	defer mutex.Unlock()

	val, found := debts[id]
	return val, found
}
