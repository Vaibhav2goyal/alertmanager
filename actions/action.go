package actions

import "alertmanager/models"

type Action interface {
	Execute(alert models.Alert)
}

var actions []Action

func TakeAction(alert models.Alert) {
	for _, action := range actions {
		action.Execute(alert)
	}
}

func RegisterAction(a Action) {
	actions = append(actions, a)
}
