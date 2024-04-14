package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

var letters = []rune(" abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789\n.,")

const wc = 342190 / 4

func Create_random_file(fp string, len int) error {
	if _, err := os.Stat(fp); errors.Is(err, os.ErrNotExist) {
		fmt.Println("File not exits")
		if _, err := os.Create(fp); err != nil {
			return err

		}
	}
	f, err := os.OpenFile(fp, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()

	if err != nil {
		return err
	}

	for i := 0; i < len; i++ {
		l := letters[rand.Intn(64)] // Len of letters
		_, err := f.WriteString(string(l))
		if err != nil {
			return err
		}
	}
	return nil
}
