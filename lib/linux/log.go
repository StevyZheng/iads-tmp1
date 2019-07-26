package linux

import (
	"bufio"
	"github.com/emirpasic/gods/lists/arraylist"
	"iads/lib/base"
	"io"
	"os"
	"strings"
)

var MsgMap = map[string]string{}

func FillMsgMap(){
	MsgMap["Initializing cgroup subsys cpuset"] = "system start"
	MsgMap["BIOS bug"] = "BIOS bug"
	MsgMap["mlx5_core"] = "mlx5 fw"
	MsgMap["Initialized ast"] = "initialized ast"
}

type LogLine struct {
	Month string
	Day int
	Hour int
	Minute int
	Second int
	Host string
	Speaker string
	Context string
	Status string
}

func NewLogLine(month string, day int, hour int, minute int, second int, host string, speaker string, context string) *LogLine {
	var result = new(LogLine)
	result.Month = month
	result.Day = day
	result.Hour = hour
	result.Minute = minute
	result.Second = second
	result.Host = host
	result.Speaker = speaker
	result.Context = context
	return result
}

func NewLine(line string) *LogLine {
	arr := base.SplitString(line, " ")
	var result = new(LogLine)
	result.Month = arr[0]
	result.Day = base.StrToInt(arr[1])
	timeArr := base.SplitString(arr[2], ":")
	result.Hour = base.StrToInt(timeArr[0])
	result.Minute = base.StrToInt(timeArr[1])
	result.Second = base.StrToInt(timeArr[2])
	result.Host = arr[3]
	result.Speaker = arr[4]
	tmp := ""
	for i := 5; i < len(arr); i++{
		tmp = tmp + " " + arr[i]
	}
	result.Context = tmp
	return result
}

type Messages struct {
	MsgList *arraylist.List
	EventList *arraylist.List
}

func NewMessages() *Messages {
	msgs := new(Messages)
	msgs.MsgList = arraylist.New()
	msgs.EventList = arraylist.New()
	return msgs
}

func (m *Messages)LoadEvent(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		lineStr, err := buf.ReadString('\n')
		if err != nil{
			if err == io.EOF{
				return nil
			}
			return err
		}
		lineStr = strings.TrimSpace(lineStr)
		for k, v := range MsgMap{
			if true == base.ContainStr(k, lineStr){
				line := NewLine(lineStr)
				line.Status = v
				m.EventList.Add(line)
			}
		}
	}
	return nil
}

func (m *Messages)LoadMessages(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil{
		panic(err)
	}
	defer f.Close()
	buf := bufio.NewReader(f)
	for {
		lineStr, err := buf.ReadString('\n')
		lineStr = strings.TrimSpace(lineStr)
		line := NewLine(lineStr)
		m.AddLine(*line)
		if err != nil{
			if err == io.EOF{
				return nil
			}
			return err
		}
	}
	return nil
}

func (m *Messages)AddLine(line LogLine) {
	m.MsgList.Add(line)
}

func (m *Messages)DelLine(index int) {
	m.MsgList.Remove(index)
}


