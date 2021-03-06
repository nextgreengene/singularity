package action

var ActionStorage = map[string]func(){
	"ExecuteExit":         ExecuteExit,
	"ExecuteStart":        ExecuteStart,
	"ExecuteDeactivate":   ExecuteDeactivate,
	"ExecuteShowPlayers":  ExecuteShowPlayers,
	"ExecuteShowStats":    ExecuteShowStats,
	"ExecuteHelp":         ExecuteHelp,
	"ExecuteCreatePlayer": ExecuteCreatePlayer,
	"ExecuteShowContext":  ExecuteShowContext,
	"ExecuteRandomPlayer": ExecuteRandomPlayer,
}

var InputActionStorage = map[string]func([]string){
	"ExecuteSelectPlayer": ExecuteSelectPlayer,
}
