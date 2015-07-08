package script

import (
	"fmt"
	"strconv"
	"strings"
)

//Operation types
const (
	_ = iota
	OpCall
	OpConditional
	OpAssignment
	OpError
)

type Operation struct {
	opType     int
	offset     int
	length     int
	opCode     byte
	assignDst  string
	assignVal  string
	condOp1    string
	condOp     string
	condOp2    string
	condDst    int
	callMethod string
	callResult string
	callMap    map[string]string
	callParams []string
	errorMsg   string
}

func (op Operation) GetMethod() string {
	if op.opType == OpCall {
		return op.callMethod
	} else {
		return ""
	}
}

type Script []Operation

type ScriptProperties struct {
	ExitTo       int
	HasExit      bool
	LoadedScript int
	LoadsScript  bool
}

func (script Script) Properties() (out ScriptProperties) {
	out = ScriptProperties{
		ExitTo:       -1,
		HasExit:      false,
		LoadedScript: -1,
		LoadsScript:  false,
	}
	for _, op := range script {
		switch op.callMethod {
		case "loadRoomWithEgo", "putActorInRoom":
			out.ExitTo, _ = strconv.Atoi(op.callMap["room"])
			out.HasExit = true
		case "startScript", "chainScript":
			out.LoadedScript, _ = strconv.Atoi(op.callMap["script"])
			out.LoadsScript = true
		}
	}
	return
}

func (script Script) Debug() string {
	out := ""
	for _, op := range script {
		out += fmt.Sprintf("[%04X] (%02x) %v\n",
			op.offset, op.opCode, op.Debug())
	}
	return out
}

func (script Script) Print() string {
	out := ""
	indent := 0
	condUntil := make([]int, 0)
	for _, op := range script {
		out += strings.Repeat("  ", indent)
		out += op.String()
		if op.opType == OpConditional && op.offset < op.condDst {
			condUntil = append(condUntil, op.condDst)
			indent++
		}
		out += "\n"
		if indent > 0 && condUntil[indent-1] == op.offset+op.length {
			condUntil = condUntil[0 : indent-1]
			indent--
			out += strings.Repeat("  ", indent)
			out += "}\n"
		}
	}
	return out
}

func (op *Operation) addNamedStringParam(paramName string, value string) {
	op.callMap[paramName] = "\"" + value + "\""
}

func (op *Operation) addNamedParam(paramName string, value int) {
	op.callMap[paramName] = fmt.Sprintf("%d", value)
}

func (op *Operation) addParam(param string) {
	op.callParams = append(op.callParams, param)
}

func (op *Operation) addResult(result string) {
	op.callResult = result
}

func (op Operation) Debug() string {
	if op.opType == OpCall {
		params := ""
		for _, param := range op.callParams {
			if params != "" {
				params += ", "
			}
			params += param
		}
		for paramName := range op.callMap {
			if params != "" {
				params += ", "
			}
			params += paramName + "=" + op.callMap[paramName]
		}
		callResult := ""
		if op.callResult != "" {
			callResult += fmt.Sprintf("%v = ", op.callResult)
		}
		return fmt.Sprintf("%v%v(%v)", callResult, op.callMethod, params)
	} else if op.opType == OpAssignment {
		return fmt.Sprintf("%v = %v", op.assignDst, op.assignVal)
	} else if op.opType == OpConditional {
		return fmt.Sprintf("unless (%v %v %v) goto %04x", op.condOp1, op.condOp, op.condOp2, op.condDst)
	} else if op.opType == OpError {
		return fmt.Sprintf("%v", op.errorMsg)
	}
	return ""
}

func (op Operation) String() string {
	if op.opType == OpConditional {
		return fmt.Sprintf("unless (%v %v %v) {", op.condOp1, op.condOp, op.condOp2)
	} else {
		return op.Debug()
	}
}
