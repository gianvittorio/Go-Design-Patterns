package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// sync.Once init() -- thread safety
// lazyness

var once sync.Once
//var instance Database
var instance *singletonDatabase

func readData(path string) (map[string]int, error) {
	ex, err := os.Executable()
	if err != nil {
	  panic(err)
	}
	exPath := filepath.Dir(ex)
  
	file, err := os.Open(exPath + path)
	if err != nil {
	  return nil, err
	}
	defer file.Close()
  
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
  
	result := map[string]int{}
	
	for scanner.Scan() {
	  k := scanner.Text()
	  scanner.Scan()
	  v, _ := strconv.Atoi(scanner.Text())
	  result[k] = v
	}
  
	return result, nil
  }

func GetSingletonDabatase() *singletonDatabase {
	once.Do(func() {
		caps, e := readData("capitals.txt")
		db := singletonDatabase{caps}

		if e == nil {
			db.capitals = caps
		}

		instance = &db
	})

	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDabatase().GetPopulation(city)
	}
	return result
}

func main() {
	db := GetSingletonDabatase()
	pop := db.GetPopulation("Seoul")
	fmt.Println("Pop of Seoul = ", pop)
}
