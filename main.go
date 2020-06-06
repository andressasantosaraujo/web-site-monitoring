package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	introduceShow()
	for {
		showMenu()
		optionChoose(getOption())
	}
}

func introduceShow() {
	fmt.Printf("Hello. Let's monitoring your websites? Choose your option *__*\n")
	fmt.Printf("ALERT! To monitor, list the sites here, in a file called sites.txt\n")
}

func showMenu() {
	fmt.Println(" 1 - Monitoring start")
	fmt.Println(" 2 - Show logs")
	fmt.Println(" 3 - Exit")
	fmt.Println("**********************************************************************")
}

func getOption() int {
	var option int
	fmt.Scan(&option)
	fmt.Println("The option ", option, "was choosen")
	fmt.Println("**********************************************************************")
	return option
}

func optionChoose(option int) {
	switch option {
	case 1:
		monitoringStart()
	case 2:
		logPrint()
	case 3:
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Unknown option :/ ")
		os.Exit(-1)
	}
}

func monitoringStart() {
	fmt.Println("Monitoring...")
	sites := getFileSite()
	for _, site := range sites {
		siteTest(site)
	}
}

func siteTest(site string) {
	resp, err := http.Get(site)
	if err != nil {
		fmt.Println("Error occurred: ", err, " :(")
	}
	if resp.StatusCode == 200 {
		setLog(site, true)
	} else {
		setLog(site, false)
	}
}

func getFileSite() []string {
	var sites []string
	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Error occurred: ", err, " :(")
	}
	reader := bufio.NewReader(file)
	for {
		site, err := reader.ReadString('\n')
		site = strings.TrimSpace(site)
		sites = append(sites, site)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return sites
}

func setLog(site string, status bool) {
	log, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error occurred: ", err, " :(")
	}
	log.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	log.Close()
}

func logPrint() {
	log, err := ioutil.ReadFile("log.txt")
	if err != nil {
		fmt.Println("Error occurred ", err, " :(")
	}
	fmt.Println(string(log))
}
