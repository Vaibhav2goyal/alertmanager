package actions

import "alertmanager/models"

type Action interface {
	Execute(alert models.Alert)
}

var actions []Action

// Take all the registered actions
func TakeAction(alert models.Alert) {
	for _, action := range actions {
		action.Execute(alert)
	}
}

// Register the actions
func RegisterAction(a Action) {
	actions = append(actions, a)
}
