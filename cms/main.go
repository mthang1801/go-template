package main

import "fmt"

type notification interface {
	importance() int
}

type directMessage struct {
	isUrgent       bool
	priorityLevel  int
	senderUsername string
	messageContent string
}

type groupMessage struct {
	priorityLevel  int
	groupName      string
	messageContent string
}

type systemAlert struct {
	messageContent string
	alertCode      string
}

func (dm directMessage) importance() int {
	if dm.isUrgent {
		return 50
	}
	return dm.priorityLevel
}

func (gm groupMessage) importance() int {
	return gm.priorityLevel
}

func (sa systemAlert) importance() int {
	return 100
}

func processNotification (n notification) (string, int) {
	switch n:= n.(type) {
		case directMessage:
      return n.senderUsername, n.importance()
    case groupMessage:
      return n.groupName, n.importance()
    case systemAlert:
      return n.alertCode, n.importance()
    default:
      return "", 0
	}
}

func main(){
	dm := directMessage{isUrgent: true, priorityLevel: 40, senderUsername: "John Doe", messageContent: "This is an urgent message"}
  gm := groupMessage{priorityLevel: 80, groupName: "Team 1", messageContent: "This is a group message"}
  sa := systemAlert{messageContent: "System Alert: High CPU Usage", alertCode: "CPU_001"}

  sender, importance := processNotification(dm)
  fmt.Printf("Direct Message from %s: Importance = %d\n", sender, importance)

  sender, importance = processNotification(gm)
  fmt.Printf("Group Message in %s: Importance = %d\n", sender, importance)

  sender, importance = processNotification(sa)
  fmt.Printf("System Alert: %s: Importance = %d\n", sender, importance)
}