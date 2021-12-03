package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"testing"
)

func almostEqual(v1, v2 float64) bool {
	return abs(v1-v2) <= 0.001
}

func getTestCases() map[int]float64 {
	bytes, err := ioutil.ReadFile("./sqrt_cases.csv")
	if err != nil {
		log.Println(err)
	}

	cases := make(map[int]float64)

	scanner := bufio.NewScanner(strings.NewReader(string(bytes)))
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), ",")
		key, err := strconv.Atoi(fields[0])
		if err != nil {
			log.Println(err)
		}
		val, err := strconv.ParseFloat(fields[1], 64)
		if err != nil {
			log.Println(err)
		}
		cases[key] = val
	}
	log.Println(cases)

	return cases
}

func TestSQRT(t *testing.T) {
	testCases := getTestCases()

	for key, val := range testCases {
		t.Run(fmt.Sprintf("%v", key), func(t *testing.T) {
			rt, err := sqrt(float64(key))
			if err != nil {
				log.Println(err)
			}
			if !almostEqual(val, rt) {
				t.Fatalf("bad value - %f", val)
			}
		})
	}
}
