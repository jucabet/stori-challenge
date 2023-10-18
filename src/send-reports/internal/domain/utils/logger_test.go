package utils

import (
	"errors"
	"testing"
)

func TestInfo(t *testing.T) {
	t.Run("Crear info log", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatal("panic occurred:", err)
			}
		}()

		Info("TestInfo", "MockInfo")
	})
}

func TestError(t *testing.T) {
	t.Run("Crear error log", func(t *testing.T) {
		defer func() {
			if err := recover(); err != nil {
				t.Fatal("panic occurred:", err)
			}
		}()

		ErrorLog("TestInfo", errors.New("MockError"))
	})
}
