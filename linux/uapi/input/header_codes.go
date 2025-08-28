package input

const (
	// INPUT_PROP_POINTER marks the device as requiring on‑screen
	// cursor control. Typical for mice, trackpads, or other
	// pointing devices.
	INPUT_PROP_POINTER PropCode = 0x00

	// INPUT_PROP_DIRECT marks the device as “direct input” where
	// touch/movement maps 1:1 to display coordinates (e.g. a
	// touchscreen).
	INPUT_PROP_DIRECT PropCode = 0x01

	// INPUT_PROP_BUTTONPAD means the touchpad has integrated
	// clickable buttons beneath its surface instead of separate
	// physical ones.
	INPUT_PROP_BUTTONPAD PropCode = 0x02

	// INPUT_PROP_SEMI_MT identifies devices that report a single
	// bounding box for multiple touches, not individual contact
	// positions. Seen in some older multi‑touch hardware.
	INPUT_PROP_SEMI_MT PropCode = 0x03

	// INPUT_PROP_TOPBUTTONPAD signals that virtual/soft buttons
	// are located along the top edge of the touchpad surface.
	INPUT_PROP_TOPBUTTONPAD PropCode = 0x04

	// INPUT_PROP_POINTING_STICK indicates a pointing stick —
	// the small joystick‑like nub often in laptop keyboards.
	INPUT_PROP_POINTING_STICK PropCode = 0x05

	// INPUT_PROP_ACCELEROMETER means the device has an internal
	// accelerometer for motion/orientation sensing.
	INPUT_PROP_ACCELEROMETER PropCode = 0x06

	// INPUT_PROP_MAX is the highest valid input property code.
	INPUT_PROP_MAX PropCode = 0x1f

	// INPUT_PROP_CNT is the total number of input properties.
	INPUT_PROP_CNT PropCode = INPUT_PROP_MAX + 1

	// EV_SYN marks a synchronization event, used to signal
	// boundaries between groups of related events.
	EV_SYN EventCode = 0x00

	// EV_KEY covers key/button events: press/release actions
	// for keyboards, mice, and game controllers.
	EV_KEY EventCode = 0x01

	// EV_REL indicates relative movement (position deltas),
	// such as mouse motion or scroll wheel steps.
	EV_REL EventCode = 0x02

	// EV_ABS reports absolute position values, e.g. touchscreen
	// coordinates or joystick axes.
	EV_ABS EventCode = 0x03

	// EV_MSC carries miscellaneous data that does not fit into
	// other event types, such as scancodes.
	EV_MSC EventCode = 0x04

	// EV_SW covers binary hardware switches, such as lid open/
	// closed or tablet mode toggles.
	EV_SW EventCode = 0x05

	// EV_LED controls device‑attached LEDs, like Caps Lock or
	// Num Lock indicators.
	EV_LED EventCode = 0x11

	// EV_SND triggers sounds from the device itself, e.g. system
	// beeps.
	EV_SND EventCode = 0x12

	// EV_REP adjusts key repeat delays and rates for held keys.
	EV_REP EventCode = 0x14

	// EV_FF controls force‑feedback effects such as rumble.
	EV_FF EventCode = 0x15

	// EV_PWR signals power management events (e.g. power button).
	EV_PWR EventCode = 0x16

	// EV_FF_STATUS reports feedback effect status, completion,
	// or errors.
	EV_FF_STATUS EventCode = 0x17

	// EV_MAX is the highest defined event type code.
	EV_MAX EventCode = 0x1f

	// EV_CNT is the total number of event types.
	EV_CNT EventCode = EV_MAX + 1

	// SYN_REPORT marks the end of a batch of events, making all
	// changes visible to user space.
	SYN_REPORT SyncCode = 0

	// SYN_CONFIG signals a change in device configuration.
	SYN_CONFIG SyncCode = 1

	// SYN_MT_REPORT groups data for a single touch contact in a
	// multi‑touch stream.
	SYN_MT_REPORT SyncCode = 2

	// SYN_DROPPED warns that one or more events were lost; user
	// space should re‑sync.
	SYN_DROPPED SyncCode = 3

	// SYN_MAX is the highest synchronization event code.
	SYN_MAX SyncCode = 0x0f

	// SYN_CNT is the total number of synchronization event codes.
	SYN_CNT SyncCode = SYN_MAX + 1

	// KEY_RESERVED is reserved and should not be assigned to any
	// physical key.
	KEY_RESERVED KeyCode = 0

	// KEY_ESC is the Escape key.
	KEY_ESC KeyCode = 1

	// KEY_1 is the '1' key.
	KEY_1 KeyCode = 2

	// KEY_2 is the '2' key.
	KEY_2 KeyCode = 3

	// KEY_3 is the '3' key.
	KEY_3 KeyCode = 4

	// KEY_4 is the '4' key.
	KEY_4 KeyCode = 5

	// KEY_5 is the '5' key.
	KEY_5 KeyCode = 6

	// KEY_6 is the '6' key.
	KEY_6 KeyCode = 7

	// KEY_7 is the '7' key.
	KEY_7 KeyCode = 8

	// KEY_8 is the '8' key.
	KEY_8 KeyCode = 9

	// KEY_9 is the '9' key.
	KEY_9 KeyCode = 10

	// KEY_0 is the '0' key.
	KEY_0 KeyCode = 11

	// KEY_MINUS is the '-' key.
	KEY_MINUS KeyCode = 12

	// KEY_EQUAL is the '=' key.
	KEY_EQUAL KeyCode = 13

	// KEY_BACKSPACE is the Backspace key.
	KEY_BACKSPACE KeyCode = 14

	// KEY_TAB is the Tab key.
	KEY_TAB KeyCode = 15

	// KEY_Q is the 'Q' key.
	KEY_Q KeyCode = 16

	// KEY_W is the 'W' key.
	KEY_W KeyCode = 17

	// KEY_E is the 'E' key.
	KEY_E KeyCode = 18

	// KEY_R is the 'R' key.
	KEY_R KeyCode = 19

	// KEY_T is the 'T' key.
	KEY_T KeyCode = 20

	// KEY_Y is the 'Y' key.
	KEY_Y KeyCode = 21

	// KEY_U is the 'U' key.
	KEY_U KeyCode = 22

	// KEY_I is the 'I' key.
	KEY_I KeyCode = 23

	// KEY_O is the 'O' key.
	KEY_O KeyCode = 24

	// KEY_P is the 'P' key.
	KEY_P KeyCode = 25

	// KEY_LEFTBRACE is the '[' key.
	KEY_LEFTBRACE KeyCode = 26

	// KEY_RIGHTBRACE is the ']' key.
	KEY_RIGHTBRACE KeyCode = 27

	// KEY_ENTER is the Enter key.
	KEY_ENTER KeyCode = 28

	// KEY_LEFTCTRL is the left Control key.
	KEY_LEFTCTRL KeyCode = 29

	// KEY_A is the 'A' key.
	KEY_A KeyCode = 30

	// KEY_S is the 'S' key.
	KEY_S KeyCode = 31

	// KEY_D is the 'D' key.
	KEY_D KeyCode = 32

	// KEY_F is the 'F' key.
	KEY_F KeyCode = 33

	// KEY_G is the 'G' key.
	KEY_G KeyCode = 34

	// KEY_H is the 'H' key.
	KEY_H KeyCode = 35

	// KEY_J is the 'J' key.
	KEY_J KeyCode = 36

	// KEY_K is the 'K' key.
	KEY_K KeyCode = 37

	// KEY_L is the 'L' key.
	KEY_L KeyCode = 38

	// KEY_SEMICOLON is the ';' key.
	KEY_SEMICOLON KeyCode = 39

	// KEY_APOSTROPHE is the ”' key.
	KEY_APOSTROPHE KeyCode = 40

	// KEY_GRAVE is the '`' key.
	KEY_GRAVE KeyCode = 41

	// KEY_LEFTSHIFT is the left Shift key.
	KEY_LEFTSHIFT KeyCode = 42

	// KEY_BACKSLASH is the '\' key.
	KEY_BACKSLASH KeyCode = 43

	// KEY_Z is the 'Z' key.
	KEY_Z KeyCode = 44

	// KEY_X is the 'X' key.
	KEY_X KeyCode = 45

	// KEY_C is the 'C' key.
	KEY_C KeyCode = 46

	// KEY_V is the 'V' key.
	KEY_V KeyCode = 47

	// KEY_B is the 'B' key.
	KEY_B KeyCode = 48

	// KEY_N is the 'N' key.
	KEY_N KeyCode = 49

	// KEY_M is the 'M' key.
	KEY_M KeyCode = 50

	// KEY_COMMA is the ',' key.
	KEY_COMMA KeyCode = 51

	// KEY_DOT is the '.' key.
	KEY_DOT KeyCode = 52

	// KEY_SLASH is the '/' key.
	KEY_SLASH KeyCode = 53

	// KEY_RIGHTSHIFT is the right Shift key.
	KEY_RIGHTSHIFT KeyCode = 54

	// KEY_KPASTERISK is the keypad '*' key.
	KEY_KPASTERISK KeyCode = 55

	// KEY_LEFTALT is the left Alt key.
	KEY_LEFTALT KeyCode = 56

	// KEY_SPACE is the Space Bar.
	KEY_SPACE KeyCode = 57

	// KEY_CAPSLOCK is the Caps Lock key.
	KEY_CAPSLOCK KeyCode = 58

	// KEY_F1 is the F1 function key.
	KEY_F1 KeyCode = 59

	// KEY_F2 is the F2 function key.
	KEY_F2 KeyCode = 60

	// KEY_F3 is the F3 function key.
	KEY_F3 KeyCode = 61

	// KEY_F4 is the F4 function key.
	KEY_F4 KeyCode = 62

	// KEY_F5 is the F5 function key.
	KEY_F5 KeyCode = 63

	// KEY_F6 is the F6 function key.
	KEY_F6 KeyCode = 64

	// KEY_F7 is the F7 function key.
	KEY_F7 KeyCode = 65

	// KEY_F8 is the F8 function key.
	KEY_F8 KeyCode = 66

	// KEY_F9 is the F9 function key.
	KEY_F9 KeyCode = 67

	// KEY_F10 is the F10 function key.
	KEY_F10 KeyCode = 68

	// KEY_NUMLOCK is the Num Lock key.
	KEY_NUMLOCK KeyCode = 69

	// KEY_SCROLLLOCK is the Scroll Lock key.
	KEY_SCROLLLOCK KeyCode = 70

	// KEY_KP7 is the keypad '7' key.
	KEY_KP7 KeyCode = 71

	// KEY_KP8 is the keypad '8' key.
	KEY_KP8 KeyCode = 72

	// KEY_KP9 is the keypad '9' key.
	KEY_KP9 KeyCode = 73

	// KEY_KPMINUS is the keypad '-' key.
	KEY_KPMINUS KeyCode = 74

	// KEY_KP4 is the keypad '4' key.
	KEY_KP4 KeyCode = 75

	// KEY_KP5 is the keypad '5' key.
	KEY_KP5 KeyCode = 76

	// KEY_KP6 is the keypad '6' key.
	KEY_KP6 KeyCode = 77

	// KEY_KPPLUS is the keypad '+' key.
	KEY_KPPLUS KeyCode = 78

	// KEY_KP1 is the keypad '1' key.
	KEY_KP1 KeyCode = 79

	// KEY_KP2 is the keypad '2' key.
	KEY_KP2 KeyCode = 80

	// KEY_KP3 is the keypad '3' key.
	KEY_KP3 KeyCode = 81

	// KEY_KP0 is the keypad '0' key.
	KEY_KP0 KeyCode = 82

	// KEY_KPDOT is the keypad '.' key.
	KEY_KPDOT KeyCode = 83

	// KEY_ZENKAKUHANKAKU is the Zenkaku/Hankaku key.
	KEY_ZENKAKUHANKAKU KeyCode = 85

	// KEY_102ND is the ISO 102nd keyboard key.
	KEY_102ND KeyCode = 86

	// KEY_F11 is the F11 function key.
	KEY_F11 KeyCode = 87

	// KEY_F12 is the F12 function key.
	KEY_F12 KeyCode = 88

	// KEY_RO is the RO key on Japanese keyboards.
	KEY_RO KeyCode = 89

	// KEY_KATAKANA is the Katakana key.
	KEY_KATAKANA KeyCode = 90

	// KEY_HIRAGANA is the Hiragana key.
	KEY_HIRAGANA KeyCode = 91

	// KEY_HENKAN is the Henkan (conversion) key.
	KEY_HENKAN KeyCode = 92

	// KEY_KATAKANAHIRAGANA toggles between Katakana and Hiragana.
	KEY_KATAKANAHIRAGANA KeyCode = 93

	// KEY_MUHENKAN is the Muhenkan (non‑conversion) key.
	KEY_MUHENKAN KeyCode = 94

	// KEY_KPJPCOMMA is the keypad Japanese comma key.
	KEY_KPJPCOMMA KeyCode = 95

	// KEY_KPENTER is the keypad Enter key.
	KEY_KPENTER KeyCode = 96

	// KEY_RIGHTCTRL is the right Control key.
	KEY_RIGHTCTRL KeyCode = 97

	// KEY_KPSLASH is the keypad slash key.
	KEY_KPSLASH KeyCode = 98

	// KEY_SYSRQ is the System Request (SysRq) key.
	KEY_SYSRQ KeyCode = 99

	// KEY_RIGHTALT is the right Alt key.
	KEY_RIGHTALT KeyCode = 100

	// KEY_LINEFEED is the Line Feed key.
	KEY_LINEFEED KeyCode = 101

	// KEY_HOME is the Home key.
	KEY_HOME KeyCode = 102

	// KEY_UP is the Up Arrow key.
	KEY_UP KeyCode = 103

	// KEY_PAGEUP is the Page Up key.
	KEY_PAGEUP KeyCode = 104

	// KEY_LEFT is the Left Arrow key.
	KEY_LEFT KeyCode = 105

	// KEY_RIGHT is the Right Arrow key.
	KEY_RIGHT KeyCode = 106

	// KEY_END is the End key.
	KEY_END KeyCode = 107

	// KEY_DOWN is the Down Arrow key.
	KEY_DOWN KeyCode = 108

	// KEY_PAGEDOWN is the Page Down key.
	KEY_PAGEDOWN KeyCode = 109

	// KEY_INSERT is the Insert key.
	KEY_INSERT KeyCode = 110

	// KEY_DELETE is the Delete key.
	KEY_DELETE KeyCode = 111

	// KEY_MACRO is the Macro key.
	KEY_MACRO KeyCode = 112

	// KEY_MUTE is the Mute key.
	KEY_MUTE KeyCode = 113

	// KEY_VOLUMEDOWN is the Volume Down key.
	KEY_VOLUMEDOWN KeyCode = 114

	// KEY_VOLUMEUP is the Volume Up key.
	KEY_VOLUMEUP KeyCode = 115

	// KEY_POWER is the System Power Down key.
	KEY_POWER KeyCode = 116

	// KEY_KPEQUAL is the keypad Equal key.
	KEY_KPEQUAL KeyCode = 117

	// KEY_KPPLUSMINUS is the keypad Plus/Minus key.
	KEY_KPPLUSMINUS KeyCode = 118

	// KEY_PAUSE is the Pause key.
	KEY_PAUSE KeyCode = 119

	// KEY_SCALE is the Compiz Scale (Expose) key.
	KEY_SCALE KeyCode = 120

	// KEY_KPCOMMA is the keypad comma key.
	KEY_KPCOMMA KeyCode = 121

	// KEY_HANGEUL is the Hangeul key on Korean keyboards.
	KEY_HANGEUL KeyCode = 122

	// KEY_HANGUEL is an alias for KEY_HANGEUL.
	KEY_HANGUEL KeyCode = KEY_HANGEUL

	// KEY_HANJA is the Hanja key on Korean keyboards.
	KEY_HANJA KeyCode = 123

	// KEY_YEN is the Yen currency key.
	KEY_YEN KeyCode = 124

	// KEY_LEFTMETA is the left Meta (Windows) key.
	KEY_LEFTMETA KeyCode = 125

	// KEY_RIGHTMETA is the right Meta (Windows) key.
	KEY_RIGHTMETA KeyCode = 126

	// KEY_COMPOSE is the Compose key.
	KEY_COMPOSE KeyCode = 127

	// KEY_STOP is the Application Control Stop key.
	KEY_STOP KeyCode = 128

	// KEY_AGAIN is the Application Control Again key.
	KEY_AGAIN KeyCode = 129

	// KEY_PROPS is the Application Control Properties key.
	KEY_PROPS KeyCode = 130

	// KEY_UNDO is the Application Control Undo key.
	KEY_UNDO KeyCode = 131

	// KEY_FRONT is the Application Control Front key.
	KEY_FRONT KeyCode = 132

	// KEY_COPY is the Application Control Copy key.
	KEY_COPY KeyCode = 133

	// KEY_OPEN is the Application Control Open key.
	KEY_OPEN KeyCode = 134

	// KEY_PASTE is the Application Control Paste key.
	KEY_PASTE KeyCode = 135

	// KEY_FIND is the Application Control Search key.
	KEY_FIND KeyCode = 136

	// KEY_CUT is the Application Control Cut key.
	KEY_CUT KeyCode = 137

	// KEY_HELP is the Application Launch Integrated Help Center key.
	KEY_HELP KeyCode = 138

	// KEY_MENU is the Application Launch Menu key.
	KEY_MENU KeyCode = 139

	// KEY_CALC is the Application Launch Calculator key.
	KEY_CALC KeyCode = 140

	// KEY_SETUP is the Application Control Setup key.
	KEY_SETUP KeyCode = 141

	// KEY_SLEEP is the System Control Sleep key.
	KEY_SLEEP KeyCode = 142

	// KEY_WAKEUP is the System Control Wake Up key.
	KEY_WAKEUP KeyCode = 143

	// KEY_FILE is the Application Launch Local Machine Browser key.
	KEY_FILE KeyCode = 144

	// KEY_SENDFILE is the Application Control Send File key.
	KEY_SENDFILE KeyCode = 145

	// KEY_DELETEFILE is the Application Control Delete File key.
	KEY_DELETEFILE KeyCode = 146

	// KEY_XFER is the Application Control Transfer key.
	KEY_XFER KeyCode = 147

	// KEY_PROG1 is the Application Launch Program 1 key.
	KEY_PROG1 KeyCode = 148

	// KEY_PROG2 is the Application Launch Program 2 key.
	KEY_PROG2 KeyCode = 149

	// KEY_WWW is the Application Launch Internet Browser key.
	KEY_WWW KeyCode = 150

	// KEY_MSDOS is the Application Launch MSDOS key.
	KEY_MSDOS KeyCode = 151

	// KEY_COFFEE is the Application Launch Terminal Lock/Screensaver key.
	KEY_COFFEE KeyCode = 152

	// KEY_SCREENLOCK is an alias for KEY_COFFEE.
	KEY_SCREENLOCK KeyCode = KEY_COFFEE

	// KEY_ROTATE_DISPLAY is the Application Control Display Orientation key.
	KEY_ROTATE_DISPLAY KeyCode = 153

	// KEY_DIRECTION is an alias for KEY_ROTATE_DISPLAY.
	KEY_DIRECTION KeyCode = KEY_ROTATE_DISPLAY

	// KEY_CYCLEWINDOWS is the Application Control Cycle Windows key.
	KEY_CYCLEWINDOWS KeyCode = 154

	// KEY_MAIL is the Application Launch Mail key.
	KEY_MAIL KeyCode = 155

	// KEY_BOOKMARKS is the Application Control Bookmarks key.
	KEY_BOOKMARKS KeyCode = 156

	// KEY_COMPUTER is the Application Launch Computer key.
	KEY_COMPUTER KeyCode = 157

	// KEY_BACK is the Application Control Back key.
	KEY_BACK KeyCode = 158

	// KEY_FORWARD is the Application Control Forward key.
	KEY_FORWARD KeyCode = 159

	// KEY_CLOSECD is the Application Control Close CD key.
	KEY_CLOSECD KeyCode = 160

	// KEY_EJECTCD is the Application Control Eject CD key.
	KEY_EJECTCD KeyCode = 161

	// KEY_EJECTCLOSECD is the Application Control Eject/Close CD key.
	KEY_EJECTCLOSECD KeyCode = 162

	// KEY_NEXTSONG is the Application Control Next Song key.
	KEY_NEXTSONG KeyCode = 163

	// KEY_PLAYPAUSE is the Application Control Play/Pause key.
	KEY_PLAYPAUSE KeyCode = 164

	// KEY_PREVIOUSSONG is the Application Control Previous Song key.
	KEY_PREVIOUSSONG KeyCode = 165

	// KEY_STOPCD is the Application Control Stop CD key.
	KEY_STOPCD KeyCode = 166

	// KEY_RECORD is the Application Control Record key.
	KEY_RECORD KeyCode = 167

	// KEY_REWIND is the Application Control Rewind key.
	KEY_REWIND KeyCode = 168

	// KEY_PHONE is the Application Launch Telephone key.
	KEY_PHONE KeyCode = 169

	// KEY_ISO is the ISO special function key.
	KEY_ISO KeyCode = 170

	// KEY_CONFIG is the Application Launch Consumer Control
	// Configuration key.
	KEY_CONFIG KeyCode = 171

	// KEY_HOMEPAGE is the Application Control Home key.
	KEY_HOMEPAGE KeyCode = 172

	// KEY_REFRESH is the Application Control Refresh key.
	KEY_REFRESH KeyCode = 173

	// KEY_EXIT is the Application Control Exit key.
	KEY_EXIT KeyCode = 174

	// KEY_MOVE is the Application Control Move key.
	KEY_MOVE KeyCode = 175

	// KEY_EDIT is the Application Control Edit key.
	KEY_EDIT KeyCode = 176

	// KEY_SCROLLUP is the Application Control Scroll Up key.
	KEY_SCROLLUP KeyCode = 177

	// KEY_SCROLLDOWN is the Application Control Scroll Down key.
	KEY_SCROLLDOWN KeyCode = 178

	// KEY_KPLEFTPAREN is the keypad '(' key.
	KEY_KPLEFTPAREN KeyCode = 179

	// KEY_KPRIGHTPAREN is the keypad ')' key.
	KEY_KPRIGHTPAREN KeyCode = 180

	// KEY_NEW is the Application Control New key.
	KEY_NEW KeyCode = 181

	// KEY_REDO is the Application Control Redo/Repeat key.
	KEY_REDO KeyCode = 182

	// KEY_F13 is the F13 function key.
	KEY_F13 KeyCode = 183

	// KEY_F14 is the F14 function key.
	KEY_F14 KeyCode = 184

	// KEY_F15 is the F15 function key.
	KEY_F15 KeyCode = 185

	// KEY_F16 is the F16 function key.
	KEY_F16 KeyCode = 186

	// KEY_F17 is the F17 function key.
	KEY_F17 KeyCode = 187

	// KEY_F18 is the F18 function key.
	KEY_F18 KeyCode = 188

	// KEY_F19 is the F19 function key.
	KEY_F19 KeyCode = 189

	// KEY_F20 is the F20 function key.
	KEY_F20 KeyCode = 190

	// KEY_F21 is the F21 function key.
	KEY_F21 KeyCode = 191

	// KEY_F22 is the F22 function key.
	KEY_F22 KeyCode = 192

	// KEY_F23 is the F23 function key.
	KEY_F23 KeyCode = 193

	// KEY_F24 is the F24 function key.
	KEY_F24 KeyCode = 194

	// KEY_PLAYCD is the Play CD key.
	KEY_PLAYCD KeyCode = 200

	// KEY_PAUSECD is the Pause CD key.
	KEY_PAUSECD KeyCode = 201

	// KEY_PROG3 is the Program 3 key.
	KEY_PROG3 KeyCode = 202

	// KEY_PROG4 is the Program 4 key.
	KEY_PROG4 KeyCode = 203

	// KEY_ALL_APPLICATIONS is the Application Control Desktop Show All
	// Applications key.
	KEY_ALL_APPLICATIONS KeyCode = 204

	// KEY_DASHBOARD is an alias for KEY_ALL_APPLICATIONS.
	KEY_DASHBOARD KeyCode = KEY_ALL_APPLICATIONS

	// KEY_SUSPEND is the System Control Suspend key.
	KEY_SUSPEND KeyCode = 205

	// KEY_CLOSE is the Application Control Close key.
	KEY_CLOSE KeyCode = 206

	// KEY_PLAY is the Play key.
	KEY_PLAY KeyCode = 207

	// KEY_FASTFORWARD is the Fast Forward key.
	KEY_FASTFORWARD KeyCode = 208

	// KEY_BASSBOOST is the Bass Boost key.
	KEY_BASSBOOST KeyCode = 209

	// KEY_PRINT is the Application Control Print key.
	KEY_PRINT KeyCode = 210

	// KEY_HP is the HP key.
	KEY_HP KeyCode = 211

	// KEY_CAMERA is the Camera key.
	KEY_CAMERA KeyCode = 212

	// KEY_SOUND is the Sound key.
	KEY_SOUND KeyCode = 213

	// KEY_QUESTION is the Question key.
	KEY_QUESTION KeyCode = 214

	// KEY_EMAIL is the Email key.
	KEY_EMAIL KeyCode = 215

	// KEY_CHAT is the Chat key.
	KEY_CHAT KeyCode = 216

	// KEY_SEARCH is the Search key.
	KEY_SEARCH KeyCode = 217

	// KEY_CONNECT is the Connect key.
	KEY_CONNECT KeyCode = 218

	// KEY_FINANCE is the Application Launch Checkbook/Finance key.
	KEY_FINANCE KeyCode = 219

	// KEY_SPORT is the Sport key.
	KEY_SPORT KeyCode = 220

	// KEY_SHOP is the Shop key.
	KEY_SHOP KeyCode = 221

	// KEY_ALTERASE is the Alternate Erase key.
	KEY_ALTERASE KeyCode = 222

	// KEY_CANCEL is the Application Control Cancel key.
	KEY_CANCEL KeyCode = 223

	// KEY_BRIGHTNESSDOWN is the Brightness Down key.
	KEY_BRIGHTNESSDOWN KeyCode = 224

	// KEY_BRIGHTNESSUP is the Brightness Up key.
	KEY_BRIGHTNESSUP KeyCode = 225

	// KEY_MEDIA is the Media key.
	KEY_MEDIA KeyCode = 226

	// KEY_SWITCHVIDEOMODE cycles between available video outputs
	// (monitor, LCD, TV-out, etc).
	KEY_SWITCHVIDEOMODE KeyCode = 227

	// KEY_KBDILLUMTOGGLE toggles the keyboard illumination on and off.
	KEY_KBDILLUMTOGGLE KeyCode = 228

	// KEY_KBDILLUMDOWN decreases the keyboard illumination.
	KEY_KBDILLUMDOWN KeyCode = 229

	// KEY_KBDILLUMUP increases the keyboard illumination.
	KEY_KBDILLUMUP KeyCode = 230

	// KEY_SEND is the Application Control Send key.
	KEY_SEND KeyCode = 231

	// KEY_REPLY is the Application Control Reply key.
	KEY_REPLY KeyCode = 232

	// KEY_FORWARDMAIL is the Application Control Forward Msg key.
	KEY_FORWARDMAIL KeyCode = 233

	// KEY_SAVE is the Application Control Save key.
	KEY_SAVE KeyCode = 234

	// KEY_DOCUMENTS is the Documents key.
	KEY_DOCUMENTS KeyCode = 235

	// KEY_BATTERY is the Battery key.
	KEY_BATTERY KeyCode = 236

	// KEY_BLUETOOTH is the Bluetooth key.
	KEY_BLUETOOTH KeyCode = 237

	// KEY_WLAN is the WLAN key.
	KEY_WLAN KeyCode = 238

	// KEY_UWB is the UWB key.
	KEY_UWB KeyCode = 239

	// KEY_UNKNOWN is the Unknown key.
	KEY_UNKNOWN KeyCode = 240

	// KEY_VIDEO_NEXT is the drive next video source key.
	KEY_VIDEO_NEXT KeyCode = 241

	// KEY_VIDEO_PREV is the drive previous video source key.
	KEY_VIDEO_PREV KeyCode = 242

	// KEY_BRIGHTNESS_CYCLE is the Brightness Cycle key; brightness
	// increases and wraps to minimum after reaching the maximum.
	KEY_BRIGHTNESS_CYCLE KeyCode = 243

	// KEY_BRIGHTNESS_AUTO sets auto brightness; manual control is disabled
	// and relies on ambient light sensors.
	KEY_BRIGHTNESS_AUTO KeyCode = 244

	// KEY_BRIGHTNESS_ZERO is an alias for KEY_BRIGHTNESS_AUTO.
	KEY_BRIGHTNESS_ZERO KeyCode = KEY_BRIGHTNESS_AUTO

	// KEY_DISPLAY_OFF sets the display device to an off state.
	KEY_DISPLAY_OFF KeyCode = 245

	// KEY_WWAN is the Wireless WAN key (LTE, UMTS, GSM, etc.).
	KEY_WWAN KeyCode = 246

	// KEY_WIMAX is an alias for KEY_WWAN.
	KEY_WIMAX KeyCode = KEY_WWAN

	// KEY_RFKILL is the master radio control key (WiFi, Bluetooth, etc.).
	KEY_RFKILL KeyCode = 247

	// KEY_MICMUTE toggles microphone mute and unmute.
	KEY_MICMUTE KeyCode = 248

	// BTN_MISC marks the start of miscellaneous button codes.
	BTN_MISC KeyCode = 0x100

	// BTN_0 is an alias for BTN_MISC representing the first
	// miscellaneous button.
	BTN_0 KeyCode = 0x100

	// BTN_1 is the second miscellaneous button.
	BTN_1 KeyCode = 0x101

	// BTN_2 is the third miscellaneous button.
	BTN_2 KeyCode = 0x102

	// BTN_3 is the fourth miscellaneous button.
	BTN_3 KeyCode = 0x103

	// BTN_4 is the fifth miscellaneous button.
	BTN_4 KeyCode = 0x104

	// BTN_5 is the sixth miscellaneous button.
	BTN_5 KeyCode = 0x105

	// BTN_6 is the seventh miscellaneous button.
	BTN_6 KeyCode = 0x106

	// BTN_7 is the eighth miscellaneous button.
	BTN_7 KeyCode = 0x107

	// BTN_8 is the ninth miscellaneous button.
	BTN_8 KeyCode = 0x108

	// BTN_9 is the tenth miscellaneous button.
	BTN_9 KeyCode = 0x109

	// BTN_MOUSE marks the start of mouse button codes.
	BTN_MOUSE KeyCode = 0x110

	// BTN_LEFT is the left mouse button.
	BTN_LEFT KeyCode = 0x110

	// BTN_RIGHT is the right mouse button.
	BTN_RIGHT KeyCode = 0x111

	// BTN_MIDDLE is the middle mouse button.
	BTN_MIDDLE KeyCode = 0x112

	// BTN_SIDE is the side mouse button.
	BTN_SIDE KeyCode = 0x113

	// BTN_EXTRA is an extra mouse button.
	BTN_EXTRA KeyCode = 0x114

	// BTN_FORWARD is the forward mouse button.
	BTN_FORWARD KeyCode = 0x115

	// BTN_BACK is the back mouse button.
	BTN_BACK KeyCode = 0x116

	// BTN_TASK is the task mouse button.
	BTN_TASK KeyCode = 0x117

	// BTN_JOYSTICK marks the start of joystick button codes.
	BTN_JOYSTICK KeyCode = 0x120

	// BTN_TRIGGER is the primary trigger button on a joystick.
	BTN_TRIGGER KeyCode = 0x120

	// BTN_THUMB is the thumb button on a joystick.
	BTN_THUMB KeyCode = 0x121

	// BTN_THUMB2 is the second thumb button on a joystick.
	BTN_THUMB2 KeyCode = 0x122

	// BTN_TOP is the top button on a joystick.
	BTN_TOP KeyCode = 0x123

	// BTN_TOP2 is the second top button on a joystick.
	BTN_TOP2 KeyCode = 0x124

	// BTN_PINKIE is the pinkie (little finger) button on a joystick.
	BTN_PINKIE KeyCode = 0x125

	// BTN_BASE is the first base button on a joystick.
	BTN_BASE KeyCode = 0x126

	// BTN_BASE2 is the second base button on a joystick.
	BTN_BASE2 KeyCode = 0x127

	// BTN_BASE3 is the third base button on a joystick.
	BTN_BASE3 KeyCode = 0x128

	// BTN_BASE4 is the fourth base button on a joystick.
	BTN_BASE4 KeyCode = 0x129

	// BTN_BASE5 is the fifth base button on a joystick.
	BTN_BASE5 KeyCode = 0x12a

	// BTN_BASE6 is the sixth base button on a joystick.
	BTN_BASE6 KeyCode = 0x12b

	// BTN_DEAD is reserved and unassigned.
	BTN_DEAD KeyCode = 0x12f

	// BTN_GAMEPAD is the first gamepad button code.
	BTN_GAMEPAD KeyCode = 0x130

	// BTN_SOUTH is the South face button on a gamepad.
	BTN_SOUTH KeyCode = 0x130

	// BTN_A is an alias for the South face button.
	BTN_A KeyCode = BTN_SOUTH

	// BTN_EAST is the East face button on a gamepad.
	BTN_EAST KeyCode = 0x131

	// BTN_B is an alias for the East face button.
	BTN_B KeyCode = BTN_EAST

	// BTN_C is the C face button on a gamepad.
	BTN_C KeyCode = 0x132

	// BTN_NORTH is the North face button on a gamepad.
	BTN_NORTH KeyCode = 0x133

	// BTN_X is an alias for the North face button.
	BTN_X KeyCode = BTN_NORTH

	// BTN_WEST is the West face button on a gamepad.
	BTN_WEST KeyCode = 0x134

	// BTN_Y is an alias for the West face button.
	BTN_Y KeyCode = BTN_WEST

	// BTN_Z is the Z face button on a gamepad.
	BTN_Z KeyCode = 0x135

	// BTN_TL is the top left shoulder button.
	BTN_TL KeyCode = 0x136

	// BTN_TR is the top right shoulder button.
	BTN_TR KeyCode = 0x137

	// BTN_TL2 is the second top left shoulder button.
	BTN_TL2 KeyCode = 0x138

	// BTN_TR2 is the second top right shoulder button.
	BTN_TR2 KeyCode = 0x139

	// BTN_SELECT is the Select button.
	BTN_SELECT KeyCode = 0x13a

	// BTN_START is the Start button.
	BTN_START KeyCode = 0x13b

	// BTN_MODE is the Mode button.
	BTN_MODE KeyCode = 0x13c

	// BTN_THUMBL is the left thumb button.
	BTN_THUMBL KeyCode = 0x13d

	// BTN_THUMBR is the right thumb button.
	BTN_THUMBR KeyCode = 0x13e

	// BTN_DIGI marks the start of digitizer tool codes.
	BTN_DIGI KeyCode = 0x140

	// BTN_TOOL_PEN is the pen tool.
	BTN_TOOL_PEN KeyCode = 0x140

	// BTN_TOOL_RUBBER is the rubber (eraser) tool.
	BTN_TOOL_RUBBER KeyCode = 0x141

	// BTN_TOOL_BRUSH is the brush tool.
	BTN_TOOL_BRUSH KeyCode = 0x142

	// BTN_TOOL_PENCIL is the pencil tool.
	BTN_TOOL_PENCIL KeyCode = 0x143

	// BTN_TOOL_AIRBRUSH is the airbrush tool.
	BTN_TOOL_AIRBRUSH KeyCode = 0x144

	// BTN_TOOL_FINGER is the finger touch tool.
	BTN_TOOL_FINGER KeyCode = 0x145

	// BTN_TOOL_MOUSE is the mouse tool.
	BTN_TOOL_MOUSE KeyCode = 0x146

	// BTN_TOOL_LENS is the lens tool.
	BTN_TOOL_LENS KeyCode = 0x147

	// BTN_TOOL_QUINTTAP detects five fingers on a trackpad.
	BTN_TOOL_QUINTTAP KeyCode = 0x148

	// BTN_STYLUS3 is the third stylus button.
	BTN_STYLUS3 KeyCode = 0x149

	// BTN_TOUCH indicates a touch event on the digitizer.
	BTN_TOUCH KeyCode = 0x14a

	// BTN_STYLUS is the primary stylus tool.
	BTN_STYLUS KeyCode = 0x14b

	// BTN_STYLUS2 is the secondary stylus tool.
	BTN_STYLUS2 KeyCode = 0x14c

	// BTN_TOOL_DOUBLETAP detects two fingers on a trackpad.
	BTN_TOOL_DOUBLETAP KeyCode = 0x14d

	// BTN_TOOL_TRIPLETAP detects three fingers on a trackpad.
	BTN_TOOL_TRIPLETAP KeyCode = 0x14e

	// BTN_TOOL_QUADTAP detects four fingers on a trackpad.
	BTN_TOOL_QUADTAP KeyCode = 0x14f

	// BTN_WHEEL is the wheel button.
	BTN_WHEEL KeyCode = 0x150

	// BTN_GEAR_DOWN is an alias for BTN_WHEEL.
	BTN_GEAR_DOWN KeyCode = BTN_WHEEL

	// BTN_GEAR_UP is the gear up button.
	BTN_GEAR_UP KeyCode = 0x151

	// KEY_OK is the OK key.
	KEY_OK KeyCode = 0x160

	// KEY_SELECT is the Select key.
	KEY_SELECT KeyCode = 0x161

	// KEY_GOTO is the Goto key.
	KEY_GOTO KeyCode = 0x162

	// KEY_CLEAR is the Clear key.
	KEY_CLEAR KeyCode = 0x163

	// KEY_POWER2 is the second Power key.
	KEY_POWER2 KeyCode = 0x164

	// KEY_OPTION is the Option key.
	KEY_OPTION KeyCode = 0x165

	// KEY_INFO is the Application Launch OEM Features/Tips/Tutorial key.
	KEY_INFO KeyCode = 0x166

	// KEY_TIME is the Time key.
	KEY_TIME KeyCode = 0x167

	// KEY_VENDOR is the Vendor key.
	KEY_VENDOR KeyCode = 0x168

	// KEY_ARCHIVE is the Archive key.
	KEY_ARCHIVE KeyCode = 0x169

	// KEY_PROGRAM is the Media Select Program Guide key.
	KEY_PROGRAM KeyCode = 0x16a

	// KEY_CHANNEL is the Channel key.
	KEY_CHANNEL KeyCode = 0x16b

	// KEY_FAVORITES is the Favorites key.
	KEY_FAVORITES KeyCode = 0x16c

	// KEY_EPG is the EPG key.
	KEY_EPG KeyCode = 0x16d

	// KEY_PVR is the Media Select Home key.
	KEY_PVR KeyCode = 0x16e

	// KEY_MHP is the MHP key.
	KEY_MHP KeyCode = 0x16f

	// KEY_LANGUAGE is the Language key.
	KEY_LANGUAGE KeyCode = 0x170

	// KEY_TITLE is the Title key.
	KEY_TITLE KeyCode = 0x171

	// KEY_SUBTITLE is the Subtitle key.
	KEY_SUBTITLE KeyCode = 0x172

	// KEY_ANGLE is the Angle key.
	KEY_ANGLE KeyCode = 0x173

	// KEY_FULL_SCREEN is the Application Control View Toggle key.
	KEY_FULL_SCREEN KeyCode = 0x174

	// KEY_ZOOM is an alias for KEY_FULL_SCREEN.
	KEY_ZOOM KeyCode = KEY_FULL_SCREEN

	// KEY_MODE is the Mode key.
	KEY_MODE KeyCode = 0x175

	// KEY_KEYBOARD is the Keyboard key.
	KEY_KEYBOARD KeyCode = 0x176

	// KEY_ASPECT_RATIO is the Aspect Ratio key.
	KEY_ASPECT_RATIO KeyCode = 0x177

	// KEY_SCREEN is an alias for KEY_ASPECT_RATIO.
	KEY_SCREEN KeyCode = KEY_ASPECT_RATIO

	// KEY_PC is the Media Select Computer key.
	KEY_PC KeyCode = 0x178

	// KEY_TV is the Media Select TV key.
	KEY_TV KeyCode = 0x179

	// KEY_TV2 is the Media Select Cable key.
	KEY_TV2 KeyCode = 0x17a

	// KEY_VCR is the Media Select VCR key.
	KEY_VCR KeyCode = 0x17b

	// KEY_VCR2 is the VCR Plus key.
	KEY_VCR2 KeyCode = 0x17c

	// KEY_SAT is the Media Select Satellite key.
	KEY_SAT KeyCode = 0x17d

	// KEY_SAT2 is the second Satellite key.
	KEY_SAT2 KeyCode = 0x17e

	// KEY_CD is the Media Select CD key.
	KEY_CD KeyCode = 0x17f

	// KEY_TAPE is the Media Select Tape key.
	KEY_TAPE KeyCode = 0x180

	// KEY_RADIO is the Radio key.
	KEY_RADIO KeyCode = 0x181

	// KEY_TUNER is the Media Select Tuner key.
	KEY_TUNER KeyCode = 0x182

	// KEY_PLAYER is the Player key.
	KEY_PLAYER KeyCode = 0x183

	// KEY_TEXT is the Text key.
	KEY_TEXT KeyCode = 0x184

	// KEY_DVD is the Media Select DVD key.
	KEY_DVD KeyCode = 0x185

	// KEY_AUX is the Aux key.
	KEY_AUX KeyCode = 0x186

	// KEY_MP3 is the MP3 key.
	KEY_MP3 KeyCode = 0x187

	// KEY_AUDIO is the Application Launch Audio Browser key.
	KEY_AUDIO KeyCode = 0x188

	// KEY_VIDEO is the Application Launch Movie Browser key.
	KEY_VIDEO KeyCode = 0x189

	// KEY_DIRECTORY is the Directory key.
	KEY_DIRECTORY KeyCode = 0x18a

	// KEY_LIST is the List key.
	KEY_LIST KeyCode = 0x18b

	// KEY_MEMO is the Media Select Messages key.
	KEY_MEMO KeyCode = 0x18c

	// KEY_CALENDAR is the Calendar key.
	KEY_CALENDAR KeyCode = 0x18d

	// KEY_RED is the Red key.
	KEY_RED KeyCode = 0x18e

	// KEY_GREEN is the Green key.
	KEY_GREEN KeyCode = 0x18f

	// KEY_YELLOW is the Yellow key.
	KEY_YELLOW KeyCode = 0x190

	// KEY_BLUE is the Blue key.
	KEY_BLUE KeyCode = 0x191

	// KEY_CHANNELUP is the Channel Increment key.
	KEY_CHANNELUP KeyCode = 0x192

	// KEY_CHANNELDOWN is the Channel Decrement key.
	KEY_CHANNELDOWN KeyCode = 0x193

	// KEY_FIRST is the First key.
	KEY_FIRST KeyCode = 0x194

	// KEY_LAST is the Recall Last key.
	KEY_LAST KeyCode = 0x195

	// KEY_AB is the AB key.
	KEY_AB KeyCode = 0x196

	// KEY_NEXT is the Next key.
	KEY_NEXT KeyCode = 0x197

	// KEY_RESTART is the Restart key.
	KEY_RESTART KeyCode = 0x198

	// KEY_SLOW is the Slow key.
	KEY_SLOW KeyCode = 0x199

	// KEY_SHUFFLE is the Shuffle key.
	KEY_SHUFFLE KeyCode = 0x19a

	// KEY_BREAK is the Break key.
	KEY_BREAK KeyCode = 0x19b

	// KEY_PREVIOUS is the Previous key.
	KEY_PREVIOUS KeyCode = 0x19c

	// KEY_DIGITS is the Digits key.
	KEY_DIGITS KeyCode = 0x19d

	// KEY_TEEN is the Teen key.
	KEY_TEEN KeyCode = 0x19e

	// KEY_TWEN is the Twen key.
	KEY_TWEN KeyCode = 0x19f

	// KEY_VIDEOPHONE is the Media Select Video Phone key.
	KEY_VIDEOPHONE KeyCode = 0x1a0

	// KEY_GAMES is the Media Select Games key.
	KEY_GAMES KeyCode = 0x1a1

	// KEY_ZOOMIN is the Application Control Zoom In key.
	KEY_ZOOMIN KeyCode = 0x1a2

	// KEY_ZOOMOUT is the Application Control Zoom Out key.
	KEY_ZOOMOUT KeyCode = 0x1a3

	// KEY_ZOOMRESET is the Application Control Zoom key.
	KEY_ZOOMRESET KeyCode = 0x1a4

	// KEY_WORDPROCESSOR is the Application Launch Word Processor key.
	KEY_WORDPROCESSOR KeyCode = 0x1a5

	// KEY_EDITOR is the Application Launch Text Editor key.
	KEY_EDITOR KeyCode = 0x1a6

	// KEY_SPREADSHEET is the Application Launch Spreadsheet key.
	KEY_SPREADSHEET KeyCode = 0x1a7

	// KEY_GRAPHICSEDITOR is the Application Launch Graphics Editor key.
	KEY_GRAPHICSEDITOR KeyCode = 0x1a8

	// KEY_PRESENTATION is the Application Launch Presentation App key.
	KEY_PRESENTATION KeyCode = 0x1a9

	// KEY_DATABASE is the Application Launch Database App key.
	KEY_DATABASE KeyCode = 0x1aa

	// KEY_NEWS is the Application Launch Newsreader key.
	KEY_NEWS KeyCode = 0x1ab

	// KEY_VOICEMAIL is the Application Launch Voicemail key.
	KEY_VOICEMAIL KeyCode = 0x1ac

	// KEY_ADDRESSBOOK is the Application Launch Contacts/Address Book key.
	KEY_ADDRESSBOOK KeyCode = 0x1ad

	// KEY_MESSENGER is the Application Launch Instant Messaging key.
	KEY_MESSENGER KeyCode = 0x1ae

	// KEY_DISPLAYTOGGLE is the Turn display (LCD) on and off key.
	KEY_DISPLAYTOGGLE KeyCode = 0x1af

	// KEY_BRIGHTNESS_TOGGLE is an alias for KEY_DISPLAYTOGGLE.
	KEY_BRIGHTNESS_TOGGLE KeyCode = KEY_DISPLAYTOGGLE

	// KEY_SPELLCHECK is the Application Launch Spell Check key.
	KEY_SPELLCHECK KeyCode = 0x1b0

	// KEY_LOGOFF is the Application Launch Logoff key.
	KEY_LOGOFF KeyCode = 0x1b1

	// KEY_DOLLAR is the dollar sign ($) key code.
	KEY_DOLLAR KeyCode = 0x1b2

	// KEY_EURO is the euro sign (€) key code.
	KEY_EURO KeyCode = 0x1b3

	// KEY_FRAMEBACK is the frame-backward transport control key code.
	KEY_FRAMEBACK KeyCode = 0x1b4

	// KEY_FRAMEFORWARD is the frame-forward transport control key code.
	KEY_FRAMEFORWARD KeyCode = 0x1b5

	// KEY_CONTEXT_MENU is the system context menu key code.
	KEY_CONTEXT_MENU KeyCode = 0x1b6

	// KEY_MEDIA_REPEAT is the media repeat transport control key code.
	KEY_MEDIA_REPEAT KeyCode = 0x1b7

	// KEY_10CHANNELSUP is the ten-channels-up key code.
	KEY_10CHANNELSUP KeyCode = 0x1b8

	// KEY_10CHANNELSDOWN is the ten-channels-down key code.
	KEY_10CHANNELSDOWN KeyCode = 0x1b9

	// KEY_IMAGES is the image browser key code.
	KEY_IMAGES KeyCode = 0x1ba

	// KEY_NOTIFICATION_CENTER is the notification center toggle key code.
	KEY_NOTIFICATION_CENTER KeyCode = 0x1bc

	// KEY_PICKUP_PHONE is the answer incoming call key code.
	KEY_PICKUP_PHONE KeyCode = 0x1bd

	// KEY_HANGUP_PHONE is the decline incoming call key code.
	KEY_HANGUP_PHONE KeyCode = 0x1be

	// KEY_LINK_PHONE is the phone sync key code.
	KEY_LINK_PHONE KeyCode = 0x1bf

	// KEY_DEL_EOL deletes text from the cursor to the end of the line.
	KEY_DEL_EOL KeyCode = 0x1c0

	// KEY_DEL_EOS deletes text from the cursor to the end of the screen.
	KEY_DEL_EOS KeyCode = 0x1c1

	// KEY_INS_LINE inserts a new line at the cursor position.
	KEY_INS_LINE KeyCode = 0x1c2

	// KEY_DEL_LINE deletes the entire current line.
	KEY_DEL_LINE KeyCode = 0x1c3

	// KEY_FN is the function (Fn) modifier key.
	KEY_FN KeyCode = 0x1d0

	// KEY_FN_ESC is the Fn+Esc key.
	KEY_FN_ESC KeyCode = 0x1d1

	// KEY_FN_F1 is the Fn+F1 key.
	KEY_FN_F1 KeyCode = 0x1d2

	// KEY_FN_F2 is the Fn+F2 key.
	KEY_FN_F2 KeyCode = 0x1d3

	// KEY_FN_F3 is the Fn+F3 key.
	KEY_FN_F3 KeyCode = 0x1d4

	// KEY_FN_F4 is the Fn+F4 key.
	KEY_FN_F4 KeyCode = 0x1d5

	// KEY_FN_F5 is the Fn+F5 key.
	KEY_FN_F5 KeyCode = 0x1d6

	// KEY_FN_F6 is the Fn+F6 key.
	KEY_FN_F6 KeyCode = 0x1d7

	// KEY_FN_F7 is the Fn+F7 key.
	KEY_FN_F7 KeyCode = 0x1d8

	// KEY_FN_F8 is the Fn+F8 key.
	KEY_FN_F8 KeyCode = 0x1d9

	// KEY_FN_F9 is the Fn+F9 key.
	KEY_FN_F9 KeyCode = 0x1da

	// KEY_FN_F10 is the Fn+F10 key.
	KEY_FN_F10 KeyCode = 0x1db

	// KEY_FN_F11 is the Fn+F11 key.
	KEY_FN_F11 KeyCode = 0x1dc

	// KEY_FN_F12 is the Fn+F12 key.
	KEY_FN_F12 KeyCode = 0x1dd

	// KEY_FN_1 is the Fn+1 key.
	KEY_FN_1 KeyCode = 0x1de

	// KEY_FN_2 is the Fn+2 key.
	KEY_FN_2 KeyCode = 0x1df

	// KEY_FN_D is the Fn+D key.
	KEY_FN_D KeyCode = 0x1e0

	// KEY_FN_E is the Fn+E key.
	KEY_FN_E KeyCode = 0x1e1

	// KEY_FN_F is the Fn+F key.
	KEY_FN_F KeyCode = 0x1e2

	// KEY_FN_S is the Fn+S key.
	KEY_FN_S KeyCode = 0x1e3

	// KEY_FN_B is the Fn+B key.
	KEY_FN_B KeyCode = 0x1e4

	// KEY_FN_RIGHT_SHIFT is the Fn+Right Shift key.
	KEY_FN_RIGHT_SHIFT KeyCode = 0x1e5

	// KEY_BRL_DOT1 is the Braille dot 1 key code.
	KEY_BRL_DOT1 KeyCode = 0x1f1

	// KEY_BRL_DOT2 is the Braille dot 2 key code.
	KEY_BRL_DOT2 KeyCode = 0x1f2

	// KEY_BRL_DOT3 is the Braille dot 3 key code.
	KEY_BRL_DOT3 KeyCode = 0x1f3

	// KEY_BRL_DOT4 is the Braille dot 4 key code.
	KEY_BRL_DOT4 KeyCode = 0x1f4

	// KEY_BRL_DOT5 is the Braille dot 5 key code.
	KEY_BRL_DOT5 KeyCode = 0x1f5

	// KEY_BRL_DOT6 is the Braille dot 6 key code.
	KEY_BRL_DOT6 KeyCode = 0x1f6

	// KEY_BRL_DOT7 is the Braille dot 7 key code.
	KEY_BRL_DOT7 KeyCode = 0x1f7

	// KEY_BRL_DOT8 is the Braille dot 8 key code.
	KEY_BRL_DOT8 KeyCode = 0x1f8

	// KEY_BRL_DOT9 is the Braille dot 9 key code.
	KEY_BRL_DOT9 KeyCode = 0x1f9

	// KEY_BRL_DOT10 is the Braille dot 10 key code.
	KEY_BRL_DOT10 KeyCode = 0x1fa

	// KEY_NUMERIC_0 is the 0 key on a phone or remote control keypad.
	KEY_NUMERIC_0 KeyCode = 0x200

	// KEY_NUMERIC_1 is the 1 key on a phone or remote control keypad.
	KEY_NUMERIC_1 KeyCode = 0x201

	// KEY_NUMERIC_2 is the 2 key on a phone or remote control keypad.
	KEY_NUMERIC_2 KeyCode = 0x202

	// KEY_NUMERIC_3 is the 3 key on a phone or remote control keypad.
	KEY_NUMERIC_3 KeyCode = 0x203

	// KEY_NUMERIC_4 is the 4 key on a phone or remote control keypad.
	KEY_NUMERIC_4 KeyCode = 0x204

	// KEY_NUMERIC_5 is the 5 key on a phone or remote control keypad.
	KEY_NUMERIC_5 KeyCode = 0x205

	// KEY_NUMERIC_6 is the 6 key on a phone or remote control keypad.
	KEY_NUMERIC_6 KeyCode = 0x206

	// KEY_NUMERIC_7 is the 7 key on a phone or remote control keypad.
	KEY_NUMERIC_7 KeyCode = 0x207

	// KEY_NUMERIC_8 is the 8 key on a phone or remote control keypad.
	KEY_NUMERIC_8 KeyCode = 0x208

	// KEY_NUMERIC_9 is the 9 key on a phone or remote control keypad.
	KEY_NUMERIC_9 KeyCode = 0x209

	// KEY_NUMERIC_STAR is the star (*) key on a phone keypad.
	KEY_NUMERIC_STAR KeyCode = 0x20a

	// KEY_NUMERIC_POUND is the pound (#) key on a phone keypad.
	KEY_NUMERIC_POUND KeyCode = 0x20b

	// KEY_NUMERIC_A is the A key on a phone keypad.
	KEY_NUMERIC_A KeyCode = 0x20c

	// KEY_NUMERIC_B is the B key on a phone keypad.
	KEY_NUMERIC_B KeyCode = 0x20d

	// KEY_NUMERIC_C is the C key on a phone keypad.
	KEY_NUMERIC_C KeyCode = 0x20e

	// KEY_NUMERIC_D is the D key on a phone keypad.
	KEY_NUMERIC_D KeyCode = 0x20f

	// KEY_CAMERA_FOCUS is the camera focus key code.
	KEY_CAMERA_FOCUS KeyCode = 0x210

	// KEY_WPS_BUTTON is the Wi-Fi Protected Setup button key code.
	KEY_WPS_BUTTON KeyCode = 0x211

	// KEY_TOUCHPAD_TOGGLE toggles the touchpad on or off.
	KEY_TOUCHPAD_TOGGLE KeyCode = 0x212

	// KEY_TOUCHPAD_ON turns the touchpad on.
	KEY_TOUCHPAD_ON KeyCode = 0x213

	// KEY_TOUCHPAD_OFF turns the touchpad off.
	KEY_TOUCHPAD_OFF KeyCode = 0x214

	// KEY_CAMERA_ZOOMIN zooms the camera in.
	KEY_CAMERA_ZOOMIN KeyCode = 0x215

	// KEY_CAMERA_ZOOMOUT zooms the camera out.
	KEY_CAMERA_ZOOMOUT KeyCode = 0x216

	// KEY_CAMERA_UP moves the camera view up.
	KEY_CAMERA_UP KeyCode = 0x217

	// KEY_CAMERA_DOWN moves the camera view down.
	KEY_CAMERA_DOWN KeyCode = 0x218

	// KEY_CAMERA_LEFT moves the camera view left.
	KEY_CAMERA_LEFT KeyCode = 0x219

	// KEY_CAMERA_RIGHT moves the camera view right.
	KEY_CAMERA_RIGHT KeyCode = 0x21a

	// KEY_ATTENDANT_ON signals attendant call on.
	KEY_ATTENDANT_ON KeyCode = 0x21b

	// KEY_ATTENDANT_OFF signals attendant call off.
	KEY_ATTENDANT_OFF KeyCode = 0x21c

	// KEY_ATTENDANT_TOGGLE toggles attendant call state.
	KEY_ATTENDANT_TOGGLE KeyCode = 0x21d

	// KEY_LIGHTS_TOGGLE toggles the reading light on or off.
	KEY_LIGHTS_TOGGLE KeyCode = 0x21e

	// BTN_DPAD_UP is the directional pad up button code.
	BTN_DPAD_UP KeyCode = 0x220

	// BTN_DPAD_DOWN is the directional pad down button code.
	BTN_DPAD_DOWN KeyCode = 0x221

	// BTN_DPAD_LEFT is the directional pad left button code.
	BTN_DPAD_LEFT KeyCode = 0x222

	// BTN_DPAD_RIGHT is the directional pad right button code.
	BTN_DPAD_RIGHT KeyCode = 0x223

	// KEY_ALS_TOGGLE toggles the ambient light sensor.
	KEY_ALS_TOGGLE KeyCode = 0x230

	// KEY_ROTATE_LOCK_TOGGLE toggles screen rotation lock.
	KEY_ROTATE_LOCK_TOGGLE KeyCode = 0x231

	// KEY_REFRESH_RATE_TOGGLE toggles display refresh rate.
	KEY_REFRESH_RATE_TOGGLE KeyCode = 0x232

	// KEY_BUTTONCONFIG is the application launch button configuration key.
	KEY_BUTTONCONFIG KeyCode = 0x240

	// KEY_TASKMANAGER is the application launch task manager key.
	KEY_TASKMANAGER KeyCode = 0x241

	// KEY_JOURNAL is the application launch log/journal key.
	KEY_JOURNAL KeyCode = 0x242

	// KEY_CONTROLPANEL is the application launch control panel key.
	KEY_CONTROLPANEL KeyCode = 0x243

	// KEY_APPSELECT is the application launch app selection key.
	KEY_APPSELECT KeyCode = 0x244

	// KEY_SCREENSAVER is the application launch screen saver key.
	KEY_SCREENSAVER KeyCode = 0x245

	// KEY_VOICECOMMAND is the voice command activation key.
	KEY_VOICECOMMAND KeyCode = 0x246

	// KEY_ASSISTANT is the context-aware assistant activation key.
	KEY_ASSISTANT KeyCode = 0x247

	// KEY_KBD_LAYOUT_NEXT selects the next keyboard layout.
	KEY_KBD_LAYOUT_NEXT KeyCode = 0x248

	// KEY_EMOJI_PICKER shows or hides the emoji picker.
	KEY_EMOJI_PICKER KeyCode = 0x249

	// KEY_DICTATE starts or stops voice dictation.
	KEY_DICTATE KeyCode = 0x24a

	// KEY_CAMERA_ACCESS_ENABLE enables programmatic camera access.
	KEY_CAMERA_ACCESS_ENABLE KeyCode = 0x24b

	// KEY_CAMERA_ACCESS_DISABLE disables programmatic camera access.
	KEY_CAMERA_ACCESS_DISABLE KeyCode = 0x24c

	// KEY_CAMERA_ACCESS_TOGGLE toggles programmatic camera access.
	KEY_CAMERA_ACCESS_TOGGLE KeyCode = 0x24d

	// KEY_ACCESSIBILITY toggles the system accessibility UI.
	KEY_ACCESSIBILITY KeyCode = 0x24e

	// KEY_DO_NOT_DISTURB toggles Do Not Disturb mode.
	KEY_DO_NOT_DISTURB KeyCode = 0x24f

	// KEY_BRIGHTNESS_MIN sets brightness to minimum.
	KEY_BRIGHTNESS_MIN KeyCode = 0x250

	// KEY_BRIGHTNESS_MAX sets brightness to maximum.
	KEY_BRIGHTNESS_MAX KeyCode = 0x251

	// KEY_KBDINPUTASSIST_PREV selects the previous input suggestion.
	KEY_KBDINPUTASSIST_PREV KeyCode = 0x260

	// KEY_KBDINPUTASSIST_NEXT selects the next input suggestion.
	KEY_KBDINPUTASSIST_NEXT KeyCode = 0x261

	// KEY_KBDINPUTASSIST_PREVGROUP moves to the previous suggestion group.
	KEY_KBDINPUTASSIST_PREVGROUP KeyCode = 0x262

	// KEY_KBDINPUTASSIST_NEXTGROUP moves to the next suggestion group.
	KEY_KBDINPUTASSIST_NEXTGROUP KeyCode = 0x263

	// KEY_KBDINPUTASSIST_ACCEPT accepts the current input suggestion.
	KEY_KBDINPUTASSIST_ACCEPT KeyCode = 0x264

	// KEY_KBDINPUTASSIST_CANCEL cancels the current input suggestion.
	KEY_KBDINPUTASSIST_CANCEL KeyCode = 0x265

	// KEY_RIGHT_UP is the diagonal up-right navigation key.
	KEY_RIGHT_UP KeyCode = 0x266

	// KEY_RIGHT_DOWN is the diagonal down-right navigation key.
	KEY_RIGHT_DOWN KeyCode = 0x267

	// KEY_LEFT_UP is the diagonal up-left navigation key.
	KEY_LEFT_UP KeyCode = 0x268

	// KEY_LEFT_DOWN is the diagonal down-left navigation key.
	KEY_LEFT_DOWN KeyCode = 0x269

	// KEY_ROOT_MENU shows the device’s root menu.
	KEY_ROOT_MENU KeyCode = 0x26a

	// KEY_MEDIA_TOP_MENU shows the top menu of media (e.g. DVD).
	KEY_MEDIA_TOP_MENU KeyCode = 0x26b

	// KEY_NUMERIC_11 is the 11 key on a phone or remote-control keypad.
	KEY_NUMERIC_11 KeyCode = 0x26c

	// KEY_NUMERIC_12 is the 12 key on a phone or remote-control keypad.
	KEY_NUMERIC_12 KeyCode = 0x26d

	// KEY_AUDIO_DESC toggles audio description for visually impaired users.
	KEY_AUDIO_DESC KeyCode = 0x26e

	// KEY_3D_MODE toggles 3D display mode.
	KEY_3D_MODE KeyCode = 0x26f

	// KEY_NEXT_FAVORITE goes to the next favorite channel.
	KEY_NEXT_FAVORITE KeyCode = 0x270

	// KEY_STOP_RECORD stops recording.
	KEY_STOP_RECORD KeyCode = 0x271

	// KEY_PAUSE_RECORD pauses recording.
	KEY_PAUSE_RECORD KeyCode = 0x272

	// KEY_VOD launches video on demand.
	KEY_VOD KeyCode = 0x273

	// KEY_UNMUTE unmutes audio.
	KEY_UNMUTE KeyCode = 0x274

	// KEY_FASTREVERSE plays content in fast reverse.
	KEY_FASTREVERSE KeyCode = 0x275

	// KEY_SLOWREVERSE plays content in slow reverse.
	KEY_SLOWREVERSE KeyCode = 0x276

	// KEY_DATA controls interactive data applications on the current
	// channel.
	KEY_DATA KeyCode = 0x277

	// KEY_ONSCREEN_KEYBOARD toggles the on-screen keyboard.
	KEY_ONSCREEN_KEYBOARD KeyCode = 0x278

	// KEY_PRIVACY_SCREEN_TOGGLE toggles the electronic privacy screen.
	KEY_PRIVACY_SCREEN_TOGGLE KeyCode = 0x279

	// KEY_SELECTIVE_SCREENSHOT captures a selected area of the screen.
	KEY_SELECTIVE_SCREENSHOT KeyCode = 0x27a

	// KEY_NEXT_ELEMENT moves focus to the next element in the user
	// interface.
	KEY_NEXT_ELEMENT KeyCode = 0x27b

	// KEY_PREVIOUS_ELEMENT moves focus to the previous element in the user
	// interface.
	KEY_PREVIOUS_ELEMENT KeyCode = 0x27c

	// KEY_AUTOPILOT_ENGAGE_TOGGLE toggles autopilot engagement.
	KEY_AUTOPILOT_ENGAGE_TOGGLE KeyCode = 0x27d

	// KEY_MARK_WAYPOINT marks the current position as a waypoint.
	KEY_MARK_WAYPOINT KeyCode = 0x27e

	// KEY_SOS sends an SOS distress signal.
	KEY_SOS KeyCode = 0x27f

	// KEY_NAV_CHART shows the navigation chart.
	KEY_NAV_CHART KeyCode = 0x280

	// KEY_FISHING_CHART shows the fishing chart.
	KEY_FISHING_CHART KeyCode = 0x281

	// KEY_SINGLE_RANGE_RADAR activates single-range radar.
	KEY_SINGLE_RANGE_RADAR KeyCode = 0x282

	// KEY_DUAL_RANGE_RADAR activates dual-range radar.
	KEY_DUAL_RANGE_RADAR KeyCode = 0x283

	// KEY_RADAR_OVERLAY toggles the radar overlay.
	KEY_RADAR_OVERLAY KeyCode = 0x284

	// KEY_TRADITIONAL_SONAR activates traditional sonar.
	KEY_TRADITIONAL_SONAR KeyCode = 0x285

	// KEY_CLEARVU_SONAR activates ClearVu down-imaging sonar.
	KEY_CLEARVU_SONAR KeyCode = 0x286

	// KEY_SIDEVU_SONAR activates SideVu side-imaging sonar.
	KEY_SIDEVU_SONAR KeyCode = 0x287

	// KEY_NAV_INFO shows navigation information.
	KEY_NAV_INFO KeyCode = 0x288

	// KEY_BRIGHTNESS_MENU opens the brightness settings menu.
	KEY_BRIGHTNESS_MENU KeyCode = 0x289

	// KEY_MACRO1 is a user-programmable macro key.
	KEY_MACRO1 KeyCode = 0x290

	// KEY_MACRO2 is a user-programmable macro key.
	KEY_MACRO2 KeyCode = 0x291

	// KEY_MACRO3 is a user-programmable macro key.
	KEY_MACRO3 KeyCode = 0x292

	// KEY_MACRO4 is a user-programmable macro key.
	KEY_MACRO4 KeyCode = 0x293

	// KEY_MACRO5 is a user-programmable macro key.
	KEY_MACRO5 KeyCode = 0x294

	// KEY_MACRO6 is a user-programmable macro key.
	KEY_MACRO6 KeyCode = 0x295

	// KEY_MACRO7 is a user-programmable macro key.
	KEY_MACRO7 KeyCode = 0x296

	// KEY_MACRO8 is a user-programmable macro key.
	KEY_MACRO8 KeyCode = 0x297

	// KEY_MACRO9 is a user-programmable macro key.
	KEY_MACRO9 KeyCode = 0x298

	// KEY_MACRO10 is a user-programmable macro key.
	KEY_MACRO10 KeyCode = 0x299

	// KEY_MACRO11 is a user-programmable macro key.
	KEY_MACRO11 KeyCode = 0x29a

	// KEY_MACRO12 is a user-programmable macro key.
	KEY_MACRO12 KeyCode = 0x29b

	// KEY_MACRO13 is a user-programmable macro key.
	KEY_MACRO13 KeyCode = 0x29c

	// KEY_MACRO14 is a user-programmable macro key.
	KEY_MACRO14 KeyCode = 0x29d

	// KEY_MACRO15 is a user-programmable macro key.
	KEY_MACRO15 KeyCode = 0x29e

	// KEY_MACRO16 is a user-programmable macro key.
	KEY_MACRO16 KeyCode = 0x29f

	// KEY_MACRO17 is a user-programmable macro key.
	KEY_MACRO17 KeyCode = 0x2a0

	// KEY_MACRO18 is a user-programmable macro key.
	KEY_MACRO18 KeyCode = 0x2a1

	// KEY_MACRO19 is a user-programmable macro key.
	KEY_MACRO19 KeyCode = 0x2a2

	// KEY_MACRO20 is a user-programmable macro key.
	KEY_MACRO20 KeyCode = 0x2a3

	// KEY_MACRO21 is a user-programmable macro key.
	KEY_MACRO21 KeyCode = 0x2a4

	// KEY_MACRO22 is a user-programmable macro key.
	KEY_MACRO22 KeyCode = 0x2a5

	// KEY_MACRO23 is a user-programmable macro key.
	KEY_MACRO23 KeyCode = 0x2a6

	// KEY_MACRO24 is a user-programmable macro key.
	KEY_MACRO24 KeyCode = 0x2a7

	// KEY_MACRO25 is a user-programmable macro key.
	KEY_MACRO25 KeyCode = 0x2a8

	// KEY_MACRO26 is a user-programmable macro key.
	KEY_MACRO26 KeyCode = 0x2a9

	// KEY_MACRO27 is a user-programmable macro key.
	KEY_MACRO27 KeyCode = 0x2aa

	// KEY_MACRO28 is a user-programmable macro key.
	KEY_MACRO28 KeyCode = 0x2ab

	// KEY_MACRO29 is a user-programmable macro key.
	KEY_MACRO29 KeyCode = 0x2ac

	// KEY_MACRO30 is a user-programmable macro key.
	KEY_MACRO30 KeyCode = 0x2ad

	// KEY_MACRO_RECORD_START starts macro recording.
	KEY_MACRO_RECORD_START KeyCode = 0x2b0

	// KEY_MACRO_RECORD_STOP stops macro recording.
	KEY_MACRO_RECORD_STOP KeyCode = 0x2b1

	// KEY_MACRO_PRESET_CYCLE cycles through macro presets.
	KEY_MACRO_PRESET_CYCLE KeyCode = 0x2b2

	// KEY_MACRO_PRESET1 selects macro preset 1.
	KEY_MACRO_PRESET1 KeyCode = 0x2b3

	// KEY_MACRO_PRESET2 selects macro preset 2.
	KEY_MACRO_PRESET2 KeyCode = 0x2b4

	// KEY_MACRO_PRESET3 selects macro preset 3.
	KEY_MACRO_PRESET3 KeyCode = 0x2b5

	// KEY_KBD_LCD_MENU1 is the first unlabeled LCD menu key.
	KEY_KBD_LCD_MENU1 KeyCode = 0x2b8

	// KEY_KBD_LCD_MENU2 is the second unlabeled LCD menu key.
	KEY_KBD_LCD_MENU2 KeyCode = 0x2b9

	// KEY_KBD_LCD_MENU3 is the third unlabeled LCD menu key.
	KEY_KBD_LCD_MENU3 KeyCode = 0x2ba

	// KEY_KBD_LCD_MENU4 is the fourth unlabeled LCD menu key.
	KEY_KBD_LCD_MENU4 KeyCode = 0x2bb

	// KEY_KBD_LCD_MENU5 is the fifth unlabeled LCD menu key.
	KEY_KBD_LCD_MENU5 KeyCode = 0x2bc

	// BTN_TRIGGER_HAPPY is the first generic extra button code.
	BTN_TRIGGER_HAPPY KeyCode = 0x2c0

	// BTN_TRIGGER_HAPPY1 is the first generic extra button code.
	BTN_TRIGGER_HAPPY1 KeyCode = 0x2c0

	// BTN_TRIGGER_HAPPY2 is the second generic extra button code.
	BTN_TRIGGER_HAPPY2 KeyCode = 0x2c1

	// BTN_TRIGGER_HAPPY3 is the third generic extra button code.
	BTN_TRIGGER_HAPPY3 KeyCode = 0x2c2

	// BTN_TRIGGER_HAPPY4 is the fourth generic extra button code.
	BTN_TRIGGER_HAPPY4 KeyCode = 0x2c3

	// BTN_TRIGGER_HAPPY5 is the fifth generic extra button code.
	BTN_TRIGGER_HAPPY5 KeyCode = 0x2c4

	// BTN_TRIGGER_HAPPY6 is the sixth generic extra button code.
	BTN_TRIGGER_HAPPY6 KeyCode = 0x2c5

	// BTN_TRIGGER_HAPPY7 is the seventh generic extra button code.
	BTN_TRIGGER_HAPPY7 KeyCode = 0x2c6

	// BTN_TRIGGER_HAPPY8 is the eighth generic extra button code.
	BTN_TRIGGER_HAPPY8 KeyCode = 0x2c7

	// BTN_TRIGGER_HAPPY9 is the ninth generic extra button code.
	BTN_TRIGGER_HAPPY9 KeyCode = 0x2c8

	// BTN_TRIGGER_HAPPY10 is the tenth generic extra button code.
	BTN_TRIGGER_HAPPY10 KeyCode = 0x2c9

	// BTN_TRIGGER_HAPPY11 is the eleventh generic extra button code.
	BTN_TRIGGER_HAPPY11 KeyCode = 0x2ca

	// BTN_TRIGGER_HAPPY12 is the twelfth generic extra button code.
	BTN_TRIGGER_HAPPY12 KeyCode = 0x2cb

	// BTN_TRIGGER_HAPPY13 is the thirteenth generic extra button code.
	BTN_TRIGGER_HAPPY13 KeyCode = 0x2cc

	// BTN_TRIGGER_HAPPY14 is the fourteenth generic extra button code.
	BTN_TRIGGER_HAPPY14 KeyCode = 0x2cd

	// BTN_TRIGGER_HAPPY15 is the fifteenth generic extra button code.
	BTN_TRIGGER_HAPPY15 KeyCode = 0x2ce

	// BTN_TRIGGER_HAPPY16 is the sixteenth generic extra button code.
	BTN_TRIGGER_HAPPY16 KeyCode = 0x2cf

	// BTN_TRIGGER_HAPPY17 is the seventeenth generic extra button code.
	BTN_TRIGGER_HAPPY17 KeyCode = 0x2d0

	// BTN_TRIGGER_HAPPY18 is the eighteenth generic extra button code.
	BTN_TRIGGER_HAPPY18 KeyCode = 0x2d1

	// BTN_TRIGGER_HAPPY19 is the nineteenth generic extra button code.
	BTN_TRIGGER_HAPPY19 KeyCode = 0x2d2

	// BTN_TRIGGER_HAPPY20 is the twentieth generic extra button code.
	BTN_TRIGGER_HAPPY20 KeyCode = 0x2d3

	// BTN_TRIGGER_HAPPY21 is the twenty-first generic extra button code.
	BTN_TRIGGER_HAPPY21 KeyCode = 0x2d4

	// BTN_TRIGGER_HAPPY22 is the twenty-second generic extra button code.
	BTN_TRIGGER_HAPPY22 KeyCode = 0x2d5

	// BTN_TRIGGER_HAPPY23 is the twenty-third generic extra button code.
	BTN_TRIGGER_HAPPY23 KeyCode = 0x2d6

	// BTN_TRIGGER_HAPPY24 is the twenty-fourth generic extra button code.
	BTN_TRIGGER_HAPPY24 KeyCode = 0x2d7

	// BTN_TRIGGER_HAPPY25 is the twenty-fifth generic extra button code.
	BTN_TRIGGER_HAPPY25 KeyCode = 0x2d8

	// BTN_TRIGGER_HAPPY26 is the twenty-sixth generic extra button code.
	BTN_TRIGGER_HAPPY26 KeyCode = 0x2d9

	// BTN_TRIGGER_HAPPY27 is the twenty-seventh generic extra button code.
	BTN_TRIGGER_HAPPY27 KeyCode = 0x2da

	// BTN_TRIGGER_HAPPY28 is the twenty-eighth generic extra button code.
	BTN_TRIGGER_HAPPY28 KeyCode = 0x2db

	// BTN_TRIGGER_HAPPY29 is the twenty-ninth generic extra button code.
	BTN_TRIGGER_HAPPY29 KeyCode = 0x2dc

	// BTN_TRIGGER_HAPPY30 is the thirtieth generic extra button code.
	BTN_TRIGGER_HAPPY30 KeyCode = 0x2dd

	// BTN_TRIGGER_HAPPY31 is the thirty-first generic extra button code.
	BTN_TRIGGER_HAPPY31 KeyCode = 0x2de

	// BTN_TRIGGER_HAPPY32 is the thirty-second generic extra button code.
	BTN_TRIGGER_HAPPY32 KeyCode = 0x2df

	// BTN_TRIGGER_HAPPY33 is the thirty-third generic extra button code.
	BTN_TRIGGER_HAPPY33 KeyCode = 0x2e0

	// BTN_TRIGGER_HAPPY34 is the thirty-fourth generic extra button code.
	BTN_TRIGGER_HAPPY34 KeyCode = 0x2e1

	// BTN_TRIGGER_HAPPY35 is the thirty-fifth generic extra button code.
	BTN_TRIGGER_HAPPY35 KeyCode = 0x2e2

	// BTN_TRIGGER_HAPPY36 is the thirty-sixth generic extra button code.
	BTN_TRIGGER_HAPPY36 KeyCode = 0x2e3

	// BTN_TRIGGER_HAPPY37 is the thirty-seventh generic extra button code.
	BTN_TRIGGER_HAPPY37 KeyCode = 0x2e4

	// BTN_TRIGGER_HAPPY38 is the thirty-eighth generic extra button code.
	BTN_TRIGGER_HAPPY38 KeyCode = 0x2e5

	// BTN_TRIGGER_HAPPY39 is the thirty-ninth generic extra button code.
	BTN_TRIGGER_HAPPY39 KeyCode = 0x2e6

	// BTN_TRIGGER_HAPPY40 is the fortieth generic extra button code.
	BTN_TRIGGER_HAPPY40 KeyCode = 0x2e7

	// KEY_MIN_INTERESTING is the lowest interesting key code.
	KEY_MIN_INTERESTING KeyCode = KEY_MUTE

	// KEY_MAX is the highest key code value.
	KEY_MAX KeyCode = 0x2ff

	// KEY_CNT is the total number of key codes.
	KEY_CNT KeyCode = KEY_MAX + 1

	// REL_X is relative movement along the X axis.
	REL_X RelativeCode = 0x00

	// REL_Y is relative movement along the Y axis.
	REL_Y RelativeCode = 0x01

	// REL_Z is relative movement along the Z axis.
	REL_Z RelativeCode = 0x02

	// REL_RX is relative rotation around the X axis.
	REL_RX RelativeCode = 0x03

	// REL_RY is relative rotation around the Y axis.
	REL_RY RelativeCode = 0x04

	// REL_RZ is relative rotation around the Z axis.
	REL_RZ RelativeCode = 0x05

	// REL_HWHEEL is relative horizontal wheel movement.
	REL_HWHEEL RelativeCode = 0x06

	// REL_DIAL is relative dial rotation.
	REL_DIAL RelativeCode = 0x07

	// REL_WHEEL is relative vertical wheel movement.
	REL_WHEEL RelativeCode = 0x08

	// REL_MISC is a miscellaneous relative axis.
	REL_MISC RelativeCode = 0x09

	// REL_RESERVED is reserved and should not be used by input drivers.
	REL_RESERVED RelativeCode = 0x0a

	// REL_WHEEL_HI_RES is the high-resolution vertical wheel axis.
	REL_WHEEL_HI_RES RelativeCode = 0x0b

	// REL_HWHEEL_HI_RES is the high-resolution horizontal wheel axis.
	REL_HWHEEL_HI_RES RelativeCode = 0x0c

	// REL_MAX is the highest relative axis code.
	REL_MAX RelativeCode = 0x0f

	// REL_CNT is the total number of relative axis codes.
	REL_CNT RelativeCode = REL_MAX + 1

	// ABS_X is the absolute position along the X axis.
	ABS_X AbsoluteCode = 0x00

	// ABS_Y is the absolute position along the Y axis.
	ABS_Y AbsoluteCode = 0x01

	// ABS_Z is the absolute position along the Z axis.
	ABS_Z AbsoluteCode = 0x02

	// ABS_RX is the absolute rotation around the X axis.
	ABS_RX AbsoluteCode = 0x03

	// ABS_RY is the absolute rotation around the Y axis.
	ABS_RY AbsoluteCode = 0x04

	// ABS_RZ is the absolute rotation around the Z axis.
	ABS_RZ AbsoluteCode = 0x05

	// ABS_THROTTLE is the throttle control axis.
	ABS_THROTTLE AbsoluteCode = 0x06

	// ABS_RUDDER is the rudder control axis.
	ABS_RUDDER AbsoluteCode = 0x07

	// ABS_WHEEL is the steering wheel control axis.
	ABS_WHEEL AbsoluteCode = 0x08

	// ABS_GAS is the gas pedal control axis.
	ABS_GAS AbsoluteCode = 0x09

	// ABS_BRAKE is the brake pedal control axis.
	ABS_BRAKE AbsoluteCode = 0x0a

	// ABS_HAT0X is the horizontal axis of hat switch 0.
	ABS_HAT0X AbsoluteCode = 0x10

	// ABS_HAT0Y is the vertical axis of hat switch 0.
	ABS_HAT0Y AbsoluteCode = 0x11

	// ABS_HAT1X is the horizontal axis of hat switch 1.
	ABS_HAT1X AbsoluteCode = 0x12

	// ABS_HAT1Y is the vertical axis of hat switch 1.
	ABS_HAT1Y AbsoluteCode = 0x13

	// ABS_HAT2X is the horizontal axis of hat switch 2.
	ABS_HAT2X AbsoluteCode = 0x14

	// ABS_HAT2Y is the vertical axis of hat switch 2.
	ABS_HAT2Y AbsoluteCode = 0x15

	// ABS_HAT3X is the horizontal axis of hat switch 3.
	ABS_HAT3X AbsoluteCode = 0x16

	// ABS_HAT3Y is the vertical axis of hat switch 3.
	ABS_HAT3Y AbsoluteCode = 0x17

	// ABS_PRESSURE is the pressure axis (e.g., stylus pressure).
	ABS_PRESSURE AbsoluteCode = 0x18

	// ABS_DISTANCE is the distance axis (e.g., stylus distance).
	ABS_DISTANCE AbsoluteCode = 0x19

	// ABS_TILT_X is the tilt angle around the X axis.
	ABS_TILT_X AbsoluteCode = 0x1a

	// ABS_TILT_Y is the tilt angle around the Y axis.
	ABS_TILT_Y AbsoluteCode = 0x1b

	// ABS_TOOL_WIDTH is the tool width axis (e.g., eraser width).
	ABS_TOOL_WIDTH AbsoluteCode = 0x1c

	// ABS_VOLUME is the absolute volume axis.
	ABS_VOLUME AbsoluteCode = 0x20

	// ABS_PROFILE is the profile axis.
	ABS_PROFILE AbsoluteCode = 0x21

	// ABS_MISC is a miscellaneous absolute axis.
	ABS_MISC AbsoluteCode = 0x28

	// ABS_RESERVED is a reserved axis code used to detect invalid events.
	ABS_RESERVED AbsoluteCode = 0x2e

	// ABS_MT_SLOT is the multi-touch slot being modified.
	ABS_MT_SLOT AbsoluteCode = 0x2f

	// ABS_MT_TOUCH_MAJOR is the major axis of the touch ellipse.
	ABS_MT_TOUCH_MAJOR AbsoluteCode = 0x30

	// ABS_MT_TOUCH_MINOR is the minor axis of the touch ellipse.
	ABS_MT_TOUCH_MINOR AbsoluteCode = 0x31

	// ABS_MT_WIDTH_MAJOR is the major axis of the approaching ellipse.
	ABS_MT_WIDTH_MAJOR AbsoluteCode = 0x32

	// ABS_MT_WIDTH_MINOR is the minor axis of the approaching ellipse.
	ABS_MT_WIDTH_MINOR AbsoluteCode = 0x33

	// ABS_MT_ORIENTATION is the orientation of the touch ellipse.
	ABS_MT_ORIENTATION AbsoluteCode = 0x34

	// ABS_MT_POSITION_X is the X coordinate of the touch position.
	ABS_MT_POSITION_X AbsoluteCode = 0x35

	// ABS_MT_POSITION_Y is the Y coordinate of the touch position.
	ABS_MT_POSITION_Y AbsoluteCode = 0x36

	// ABS_MT_TOOL_TYPE is the type of tool in contact
	// (e.g., finger or stylus).
	ABS_MT_TOOL_TYPE AbsoluteCode = 0x37

	// ABS_MT_BLOB_ID groups packets into a single blob.
	ABS_MT_BLOB_ID AbsoluteCode = 0x38

	// ABS_MT_TRACKING_ID is a unique ID for touch contacts.
	ABS_MT_TRACKING_ID AbsoluteCode = 0x39

	// ABS_MT_PRESSURE is the pressure of the touch.
	ABS_MT_PRESSURE AbsoluteCode = 0x3a

	// ABS_MT_DISTANCE is the hover distance for touch.
	ABS_MT_DISTANCE AbsoluteCode = 0x3b

	// ABS_MT_TOOL_X is the X coordinate of the tool position.
	ABS_MT_TOOL_X AbsoluteCode = 0x3c

	// ABS_MT_TOOL_Y is the Y coordinate of the tool position.
	ABS_MT_TOOL_Y AbsoluteCode = 0x3d

	// ABS_MAX is the highest absolute axis code.
	ABS_MAX AbsoluteCode = 0x3f

	// ABS_CNT is the total number of absolute axis codes.
	ABS_CNT AbsoluteCode = ABS_MAX + 1

	// SW_LID indicates the lid is closed.
	SW_LID SwitchCode = 0x00

	// SW_TABLET_MODE indicates tablet mode is active.
	SW_TABLET_MODE SwitchCode = 0x01

	// SW_HEADPHONE_INSERT indicates headphones are inserted.
	SW_HEADPHONE_INSERT SwitchCode = 0x02

	// SW_RFKILL_ALL is the RF kill master switch (radio enabled).
	SW_RFKILL_ALL SwitchCode = 0x03

	// SW_RADIO is a deprecated alias for SW_RFKILL_ALL.
	SW_RADIO SwitchCode = SW_RFKILL_ALL

	// SW_MICROPHONE_INSERT indicates a microphone is inserted.
	SW_MICROPHONE_INSERT SwitchCode = 0x04

	// SW_DOCK indicates the device is docked.
	SW_DOCK SwitchCode = 0x05

	// SW_LINEOUT_INSERT indicates a line-out jack is connected.
	SW_LINEOUT_INSERT SwitchCode = 0x06

	// SW_JACK_PHYSICAL_INSERT indicates a mechanical jack is engaged.
	SW_JACK_PHYSICAL_INSERT SwitchCode = 0x07

	// SW_VIDEOOUT_INSERT indicates a video-out connector is attached.
	SW_VIDEOOUT_INSERT SwitchCode = 0x08

	// SW_CAMERA_LENS_COVER indicates the camera lens cover is down.
	SW_CAMERA_LENS_COVER SwitchCode = 0x09

	// SW_KEYPAD_SLIDE indicates the keypad is slid out.
	SW_KEYPAD_SLIDE SwitchCode = 0x0a

	// SW_FRONT_PROXIMITY indicates the front proximity sensor is active.
	SW_FRONT_PROXIMITY SwitchCode = 0x0b

	// SW_ROTATE_LOCK indicates screen rotation is locked.
	SW_ROTATE_LOCK SwitchCode = 0x0c

	// SW_LINEIN_INSERT indicates a line-in jack is connected.
	SW_LINEIN_INSERT SwitchCode = 0x0d

	// SW_MUTE_DEVICE indicates the device is muted.
	SW_MUTE_DEVICE SwitchCode = 0x0e

	// SW_PEN_INSERTED indicates a pen is inserted.
	SW_PEN_INSERTED SwitchCode = 0x0f

	// SW_MACHINE_COVER indicates the machine cover is closed.
	SW_MACHINE_COVER SwitchCode = 0x10

	// SW_USB_INSERT indicates a USB audio device is connected.
	SW_USB_INSERT SwitchCode = 0x11

	// SW_MAX is the highest switch event code.
	SW_MAX SwitchCode = 0x11

	// SW_CNT is the total number of switch event codes.
	SW_CNT SwitchCode = SW_MAX + 1

	// MSC_SERIAL is a serial event.
	MSC_SERIAL MiscCode = 0x00

	// MSC_PULSELED is an LED pulse event.
	MSC_PULSELED MiscCode = 0x01

	// MSC_GESTURE is a gesture event.
	MSC_GESTURE MiscCode = 0x02

	// MSC_RAW is a raw data event.
	MSC_RAW MiscCode = 0x03

	// MSC_SCAN is a scan code event.
	MSC_SCAN MiscCode = 0x04

	// MSC_TIMESTAMP is a timestamp event.
	MSC_TIMESTAMP MiscCode = 0x05

	// MSC_MAX is the highest miscellaneous event code.
	MSC_MAX MiscCode = 0x07

	// MSC_CNT is the total number of miscellaneous event codes.
	MSC_CNT MiscCode = MSC_MAX + 1

	// LED_NUML is the Num Lock LED.
	LED_NUML LEDCode = 0x00

	// LED_CAPSL is the Caps Lock LED.
	LED_CAPSL LEDCode = 0x01

	// LED_SCROLLL is the Scroll Lock LED.
	LED_SCROLLL LEDCode = 0x02

	// LED_COMPOSE is the Compose LED.
	LED_COMPOSE LEDCode = 0x03

	// LED_KANA is the Kana (input mode) LED.
	LED_KANA LEDCode = 0x04

	// LED_SLEEP is the Sleep state LED.
	LED_SLEEP LEDCode = 0x05

	// LED_SUSPEND is the Suspend state LED.
	LED_SUSPEND LEDCode = 0x06

	// LED_MUTE is the Mute state LED.
	LED_MUTE LEDCode = 0x07

	// LED_MISC is the miscellaneous LED.
	LED_MISC LEDCode = 0x08

	// LED_MAIL is the Mail notification LED.
	LED_MAIL LEDCode = 0x09

	// LED_CHARGING is the Charging state LED.
	LED_CHARGING LEDCode = 0x0a

	// LED_MAX is the highest LED code.
	LED_MAX LEDCode = 0x0f

	// LED_CNT is the total number of LED codes.
	LED_CNT LEDCode = LED_MAX + 1

	// REP_DELAY is the autorepeat delay value.
	REP_DELAY RepeatCode = 0x00

	// REP_PERIOD is the autorepeat period value.
	REP_PERIOD RepeatCode = 0x01

	// REP_MAX is the highest autorepeat index.
	REP_MAX RepeatCode = 0x01

	// REP_CNT is the total number of autorepeat values.
	REP_CNT RepeatCode = REP_MAX + 1

	// SND_CLICK is the click sound code.
	SND_CLICK SoundCode = 0x00

	// SND_BELL is the bell sound code.
	SND_BELL SoundCode = 0x01

	// SND_TONE is the tone sound code.
	SND_TONE SoundCode = 0x02

	// SND_MAX is the highest sound code.
	SND_MAX SoundCode = 0x07

	// SND_CNT is the total number of sound codes.
	SND_CNT SoundCode = SND_MAX + 1
)
