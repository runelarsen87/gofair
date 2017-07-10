package gofair

import (
	"gofair/streaming"
	"os"
	"log"
	"bufio"
	"encoding/json"
)

func (h *Historical) ParseHistoricalData(directory string, listener streaming.Listener)(error) {
	file, err := os.Open(directory)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		t := scanner.Text()

		var mc streaming.MarketChangeMessage
		err := json.Unmarshal(scanner.Bytes(), &mc)
		if err != nil {
			log.Fatal(err, t)
		}

		listener.OnData(mc)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return nil
}