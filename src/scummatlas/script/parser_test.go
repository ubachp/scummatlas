package script

import "testing"
import "io/ioutil"

func readScriptOrDie(filename string, t *testing.T) []byte {
	data, err := ioutil.ReadFile("../testdata/scripts/" + filename + ".dump")
	if err != nil {
		t.Errorf("Error reading the file")
	}
	if data[8] >= 200 && data[8] < 230 {
		return data[9:]
	} else {
		return data
	}
}

func checkScriptLengthAndOpcodes(file string, expectedOpcodes []byte, t *testing.T) {
	data := readScriptOrDie(file, t)
	script := ParseScriptBlock(data)
	if len(script) != len(expectedOpcodes) {
		t.Errorf("Length mismatch, got %d and expected %d",
			len(script), len(expectedOpcodes))
	}
	for i, _ := range script {
		if len(script) <= i || len(expectedOpcodes) <= i {
			return
		}
		if script[i].opCode != expectedOpcodes[i] {
			t.Errorf("Expecting opcode %x in position %d, but got %x",
				expectedOpcodes[i], i, script[i].opCode)
		}
	}
}

func TestRoomScript1(t *testing.T) {
	checkScriptLengthAndOpcodes("monkey1_11_200",
		[]byte{0x40, 0x1A, 0x05, 0x5D, 0x2E, 0x1C, 0x2E, 0x2A, 0x80, 0x68, 0x28, 0x1C, 0x2E, 0x0A, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0xAC, 0x1A, 0x2A, 0x80, 0xA8, 0x2A, 0x80, 0x68, 0x28, 0x48, 0x0A, 0x80, 0x68, 0x28, 0x48, 0x28, 0x1A, 0x0A, 0x33, 0x80, 0x80, 0x80, 0x33, 0x07, 0x5D, 0xD8, 0x00, 0xAE, 0xD8, 0x00, 0xAE, 0xD8, 0x00, 0xAE, 0x18, 0xD8, 0x00, 0xAE, 0xD8, 0x00, 0xC0}, t)

	checkScriptLengthAndOpcodes("monkey2_11_202",
		[]byte{0x1A, 0x80, 0xE8, 0x28, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey2_11_200",
		[]byte{
			0x13, 0x11, 0x2D, 0x01, 0x2A}, t)

	checkScriptLengthAndOpcodes("monkey2_11_210",
		[]byte{0x40, 0x93, 0x91, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0xC0, 0x93, 0x2A, 0x24}, t)

	checkScriptLengthAndOpcodes("monkey2_11_209",
		[]byte{0x1A, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey2_11_208",
		[]byte{0x1A, 0x11, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x1C, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80}, t)

	checkScriptLengthAndOpcodes("monkey2_11_203",
		[]byte{0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0xAC, 0xAC, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x48, 0x18, 0x18, 0x2A, 0x80, 0x68, 0x28, 0x18}, t)
}
