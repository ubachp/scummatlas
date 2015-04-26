package scummatlas

var opCodesNames = map[byte]string{
	0x00: "stopObjectCode",
	0x01: "putActor",
	0x02: "startMusic",
	0x03: "getActorRoom",
	0x04: "isGreaterEqual",
	0x05: "drawObject",
	0x06: "getActorElevation",
	0x07: "setState",
	0x08: "isNotEqual",
	0x09: "faceActor",
	0x0a: "startScript",
	0x0b: "getVerbEntryPoint",
	0x0c: "resourceRoutines",
	0x0d: "walkActorToActor",
	0x0e: "putActorAtObject",
	0x0f: "getObjectState",
	0x10: "getObjectOwner",
	0x11: "animateActor",
	0x12: "panCameraTo",
	0x13: "actorOps",
	0x14: "print",
	0x15: "actorFromPos",
	0x16: "getRandomNumber",
	0x17: "and",
	0x18: "jumpRelative",
	0x19: "doSentence",
	0x1a: "move",
	0x1b: "multiply",
	0x1c: "startSound",
	0x1d: "ifClassOfIs",
	0x1e: "walkActorTo",
	0x1f: "isActorInBox",
	0x20: "stopMusic",
	0x22: "getAnimCounter",
	0x23: "getActorY",
	0x24: "loadRoomWithEgo",
	0x25: "pickupObject",
	0x26: "setVarRange",
	0x27: "stringOps",
	0x28: "equalZero",
	0x29: "setOwnerOf",
	0x2b: "delayVariable",
	0x2c: "cursorCommand",
	0x2d: "putActorInRoom",
	0x2e: "delay",
	0x2f: "ifNotState",
	0x30: "matrixOp",
	0x31: "getInventoryCount",
	0x32: "setCameraAt",
	0x33: "roomOps",
	0x34: "getDist",
	0x35: "findObject",
	0x36: "walkActorToObject",
	0x37: "startObject",
	0x38: "lessOrEqual",
	0x3a: "subtract",
	0x3b: "getActorScale",
	0x3c: "stopSound",
	0x3d: "findInventory",
	0x3f: "drawBox",
	0x40: "cutscene",
	0x42: "chainScript",
	0x43: "getActorX",
	0x44: "isLess",
	0x46: "increment",
	0x48: "isEqual",
	0x4c: "soundKludge",
	0x50: "pickupObject",
	0x52: "actorFollowCamera",
	0x54: "setObjectName",
	0x56: "getActorMoving",
	0x57: "or",
	0x58: "override",
	0x5a: "add",
	0x5b: "divide",
	0x5c: "oldRoomEffect",
	0x5d: "actorSetClass",
	0x60: "freezeScripts",
	0x62: "stopScript",
	0x63: "getActorFacing",
	0x66: "getClosestObjActor",
	0x67: "getStringWidth",
	0x68: "getScriptRunning",
	0x6b: "debug",
	0x6c: "getActorWidth",
	0x6e: "stopObjectScript",
	0x70: "lights",
	0x71: "getActorCostume",
	0x72: "loadRoom",
	0x78: "isGreater",
	0x7a: "verbOps",
	0x7b: "getActorWalkBox",
	0x7c: "isSoundRunning",
	0x7e: "walkActorTo",
	0x80: "breakHere",
	0x98: "systemOps",
	0xa0: "stopObjectCode",
	0xa7: "dummy",
	0xa8: "notEqualZero",
	0xab: "saveRestoreVerbs",
	0xac: "expression",
	0xae: "wait",
	0xc0: "endCutScene",
	0xc6: "decrement",
	0xcc: "pseudoRoom",
	0xd8: "printEgo",

	//from ScummVM sourcecode
	0x2a: "startScript",
	0x59: "doSentence",
	0x61: "putActor",
	0x64: "loadRoomWithEgo",
	0x69: "setOwnerOf",
	0x6a: "startScript",
	0x6d: "putActorInRoom",
	0x74: "getDist",
	0x77: "startObject",
	0x85: "drawObject",
	0x8c: "resourceRoutines",
	0x8f: "getObjectState",
	0x91: "animateActor",
	0x93: "getInventoryCount",
	0x9a: "move",
	0xa3: "getActorY",
	0xad: "putActorInRoom",
	0xba: "subtract",
	0xc3: "getActorX",
	0xc8: "isEqual",
	0xd6: "getActorMoving",
	0xe1: "putActor",
	0xff: "drawBox",
}

var cursorCommands = map[byte]string{
	0x01: "cursorOn",
	0x02: "cursorOff",
	0x03: "userInputOn",
	0x04: "userInputOff",
	0x05: "softCursorOn",
	0x06: "softCursorOff",
	0x07: "softUserInputOn",
	0x08: "softUserInputOff",
}

var actorOps = map[byte]string{
	0x00: "dummy",
	0x01: "costume",
	0x02: "step_dist",
	0x03: "sound",
	0x04: "walk_animation",
	0x05: "talk_animation",
	0x06: "stand_animation",
	0x07: "animation",
	0x08: "default",
	0x09: "elevation",
	0x0a: "animation_default",
	0x0b: "palette",
	0x0c: "talk_color",
	0x0d: "actor_name",
	0x0e: "init_animation",
	0x10: "actor_width",
	0x11: "actor_scale",
	0x12: "never_zclip",
	0x13: "always_zclip",
	0x14: "ignore_boxes",
	0x15: "follow_boxes",
	0x16: "animation_speed",
	0x17: "shadow",
}

var resourceRoutines = map[byte]string{
	0x01: "load_script",
	0x02: "load_sound",
	0x03: "load_costume",
	0x04: "load_room",
	0x05: "nuke_script",
	0x06: "nuke_sound",
	0x07: "nuke_costume",
	0x08: "nuke_room",
	0x09: "lock_script",
	0x0a: "lock_sound",
	0x0b: "lock_costume",
	0x0c: "lock_room",
	0x0d: "unlock_script",
	0x0e: "unlock_sound",
	0x0f: "unlock_costume",
	0x10: "unlock_room",
	0x11: "clear_heap",
	0x12: "load_charset",
	0x13: "nuke_charset",
	0x14: "load_object",
}

var varNames = map[byte]string{
	0:  "KEYPRESS",
	1:  "EGO",
	2:  "CAMERA_POS_X",
	3:  "HAVE_MSG",
	4:  "ROOM",
	5:  "OVERRIDE",
	6:  "MACHINE_SPEED",
	7:  "ME",
	8:  "NUM_ACTOR",
	9:  "CURRENT_LIGHTS",
	10: "CURRENTDRIVE",
	11: "TMR_1",
	12: "TMR_2",
	13: "TMR_3",
	14: "MUSIC_TIMER",
	15: "ACTOR_RANGE_MIN",
	16: "ACTOR_RANGE_MAX",
	17: "CAMERA_MIN_X",
	18: "CAMERA_MAX_X",
	19: "TIMER_NEXT",
	20: "VIRT_MOUSE_X",
	21: "VIRT_MOUSE_Y",
	22: "ROOM_RESOURCE",
	23: "LAST_SOUND",
	24: "CUTSCENEEXIT_KEY",
	25: "TALK_ACTOR",
	26: "CAMERA_FAST_X",
	27: "SCROLL_SCRIPT",
	28: "ENTRY_SCRIPT",
	29: "ENTRY_SCRIPT2",
	30: "EXIT_SCRIPT",
	31: "EXIT_SCRIPT2",
	32: "VERB_SCRIPT",
	33: "SENTENCE_SCRIPT",
	34: "INVENTORY_SCRIPT",
	35: "CUTSCENE_START_SCRIPT",
	36: "CUTSCENE_END_SCRIPT",
	37: "CHARINC",
	38: "WALKTO_OBJ",
	39: "DEBUGMODE",
	40: "HEAPSPACE",
	42: "RESTART_KEY",
	43: "PAUSE_KEY",
	44: "MOUSE_X",
	45: "MOUSE_Y",
	46: "TIMER",
	47: "TIMER_TOTAL",
	48: "SOUNDCARD",
	49: "VIDEOMODE",
	50: "MAINMENU_KEY",
	51: "FIXEDDISK",
	52: "CURSORSTATE",
	53: "USERPUT",
	54: "V5_TALK_STRING_Y",
	56: "SOUNDRESULT",
	57: "TALKSTOP_KEY",
	59: "FADE_DELAY",
	60: "NOSUBTITLES",
	64: "SOUNDPARAM",
	65: "SOUNDPARAM2",
	66: "SOUNDPARAM3",
	67: "INPUTMODE",
	68: "MEMORY_PERFORMANCE",
	69: "VIDEO_PERFORMANCE",
	70: "ROOM_FLAG",
	71: "GAME_LOADED",
	72: "NEW_ROOM",
}
