package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type JsonFile struct {
	fileName string
	prefix   string
	indent   string
}

type Schedule struct {
	JsonFile
	ScheduleMap map[string]string
}

// Load from json file
func (s *Schedule) Load() {
	b, _ := ioutil.ReadFile(s.fileName)
	json.Unmarshal(b, &s.ScheduleMap)
}

// Show the map
func (s *Schedule) Show() {
	fmt.Printf("%s\n", "进度表")
	fmt.Println("--------------------------------")
	i := 0
	for k, v := range s.ScheduleMap {
		i += 1
		fmt.Printf("%d %-10s %s\n", i, k, v)
	}
}

func (s *Schedule) Update(name, season, episode string) {
	s.ScheduleMap[name] = "S" + season + "E" + episode
}

func (s *Schedule) Delete(name string) {
	delete(s.ScheduleMap, name)
}

func (s *Schedule) Save() {
	b, _ := json.MarshalIndent(&s.ScheduleMap, s.prefix, s.indent)
	ioutil.WriteFile(s.fileName, b, 0777)

}

func initSchedule() *Schedule {

	var s = new(Schedule)
	s.ScheduleMap = map[string]string{}
	s.fileName = "schedule.json"
	s.prefix = ""
	s.indent = "    "
	s.Load()
	return s
}

func Handle(s *Schedule) {
	for {
		var input, name, season, episode string
		s.Show()
		fmt.Println()
		fmt.Println("[u] update  [d] delete  [q] quit")
		fmt.Println()
		switch fmt.Scanln(&input); input {
		case "q":
			s.Save()
			os.Exit(0)
		case "u":
			fmt.Print("Input Name, season, episode :")
			fmt.Scanln(&name, &season, &episode)
			if name == "q" {
				continue
			}
			s.Update(name, season, episode)
		case "d":
			fmt.Print("输入要删除的剧名：")
			fmt.Scanln(&name)
			if name == "q" {
				continue
			}
			s.Delete(name)
		default:
			log.Println("输入有误！")
		}
	}
}

func main() {
	s := initSchedule()
	Handle(s)
}
