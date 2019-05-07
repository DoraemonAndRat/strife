package strife

import (
	"github.com/veandco/go-sdl2/sdl"
)

// A mapping of _all_ the SDL keys.
const (
	// mapping to sdl.K_UNKNOWN
	UNKNOWN = sdl.K_UNKNOWN // "" (no name, empty string)

	KEY_RETURN     = sdl.K_RETURN     // "Return" (the Enter key (main keyboard))
	KEY_ESCAPE     = sdl.K_ESCAPE     // "Escape" (the Esc key)
	KEY_BACKSPACE  = sdl.K_BACKSPACE  // "Backspace"
	KEY_TAB        = sdl.K_TAB        // "Tab" (the Tab key)
	KEY_SPACE      = sdl.K_SPACE      // "Space" (the Space Bar key(s))
	KEY_EXCLAIM    = sdl.K_EXCLAIM    // "!"
	KEY_QUOTEDBL   = sdl.K_QUOTEDBL   // """
	KEY_HASH       = sdl.K_HASH       // "#"
	KEY_PERCENT    = sdl.K_PERCENT    // "%"
	KEY_DOLLAR     = sdl.K_DOLLAR     // "$"
	KEY_AMPERSAND  = sdl.K_AMPERSAND  // "&"
	KEY_QUOTE      = sdl.K_QUOTE      // "'"
	KEY_LEFTPAREN  = sdl.K_LEFTPAREN  // "("
	KEY_RIGHTPAREN = sdl.K_RIGHTPAREN // ")"
	KEY_ASTERISK   = sdl.K_ASTERISK   // "*"
	KEY_PLUS       = sdl.K_PLUS       // "+"
	KEY_COMMA      = sdl.K_COMMA      // ","
	KEY_MINUS      = sdl.K_MINUS      // "-"
	KEY_PERIOD     = sdl.K_PERIOD     // "."
	KEY_SLASH      = sdl.K_SLASH      // "/"
	KEY_0          = sdl.K_0          // "0"
	KEY_1          = sdl.K_1          // "1"
	KEY_2          = sdl.K_2          // "2"
	KEY_3          = sdl.K_3          // "3"
	KEY_4          = sdl.K_4          // "4"
	KEY_5          = sdl.K_5          // "5"
	KEY_6          = sdl.K_6          // "6"
	KEY_7          = sdl.K_7          // "7"
	KEY_8          = sdl.K_8          // "8"
	KEY_9          = sdl.K_9          // "9"
	KEY_COLON      = sdl.K_COLON      // ":"
	KEY_SEMICOLON  = sdl.K_SEMICOLON  // ";"
	KEY_LESS       = sdl.K_LESS       // "<"
	KEY_EQUALS     = sdl.K_EQUALS     // "="
	KEY_GREATER    = sdl.K_GREATER    // ">"
	KEY_QUESTION   = sdl.K_QUESTION   // "?"
	KEY_AT         = sdl.K_AT         // "@"

	KEY_LEFTBRACKET  = sdl.K_LEFTBRACKET  // "["
	KEY_BACKSLASH    = sdl.K_BACKSLASH    // "\"
	KEY_RIGHTBRACKET = sdl.K_RIGHTBRACKET // "]"
	KEY_CARET        = sdl.K_CARET        // "^"
	KEY_UNDERSCORE   = sdl.K_UNDERSCORE   // "_"
	KEY_BACKQUOTE    = sdl.K_BACKQUOTE    // "`"
	KEY_A            = sdl.K_a            // "A"
	KEY_B            = sdl.K_b            // "B"
	KEY_C            = sdl.K_c            // "C"
	KEY_D            = sdl.K_d            // "D"
	KEY_E            = sdl.K_e            // "E"
	KEY_F            = sdl.K_f            // "F"
	KEY_G            = sdl.K_g            // "G"
	KEY_H            = sdl.K_h            // "H"
	KEY_I            = sdl.K_i            // "I"
	KEY_J            = sdl.K_j            // "J"
	KEY_K            = sdl.K_k            // "K"
	KEY_L            = sdl.K_l            // "L"
	KEY_M            = sdl.K_m            // "M"
	KEY_N            = sdl.K_n            // "N"
	KEY_O            = sdl.K_o            // "O"
	KEY_P            = sdl.K_p            // "P"
	KEY_Q            = sdl.K_q            // "Q"
	KEY_R            = sdl.K_r            // "R"
	KEY_S            = sdl.K_s            // "S"
	KEY_T            = sdl.K_t            // "T"
	KEY_U            = sdl.K_u            // "U"
	KEY_V            = sdl.K_v            // "V"
	KEY_W            = sdl.K_w            // "W"
	KEY_X            = sdl.K_x            // "X"
	KEY_Y            = sdl.K_y            // "Y"
	KEY_Z            = sdl.K_z            // "Z"

	KEY_CAPSLOCK = sdl.K_CAPSLOCK // "CapsLock"

	KEY_F1  = sdl.K_F1  // "F1"
	KEY_F2  = sdl.K_F2  // "F2"
	KEY_F3  = sdl.K_F3  // "F3"
	KEY_F4  = sdl.K_F4  // "F4"
	KEY_F5  = sdl.K_F5  // "F5"
	KEY_F6  = sdl.K_F6  // "F6"
	KEY_F7  = sdl.K_F7  // "F7"
	KEY_F8  = sdl.K_F8  // "F8"
	KEY_F9  = sdl.K_F9  // "F9"
	KEY_F10 = sdl.K_F10 // "F10"
	KEY_F11 = sdl.K_F11 // "F11"
	KEY_F12 = sdl.K_F12 // "F12"

	KEY_PRINTSCREEN = sdl.K_PRINTSCREEN // "PrintScreen"
	KEY_SCROLLLOCK  = sdl.K_SCROLLLOCK  // "ScrollLock"
	KEY_PAUSE       = sdl.K_PAUSE       // "Pause" (the Pause / Break key)
	KEY_INSERT      = sdl.K_INSERT      // "Insert" (insert on PC, help on some Mac keyboards (but does send code 73, not 117))
	KEY_HOME        = sdl.K_HOME        // "Home"
	KEY_PAGEUP      = sdl.K_PAGEUP      // "PageUp"
	KEY_DELETE      = sdl.K_DELETE      // "Delete"
	KEY_END         = sdl.K_END         // "End"
	KEY_PAGEDOWN    = sdl.K_PAGEDOWN    // "PageDown"
	KEY_RIGHT       = sdl.K_RIGHT       // "Right" (the Right arrow key (navigation keypad))
	KEY_LEFT        = sdl.K_LEFT        // "Left" (the Left arrow key (navigation keypad))
	KEY_DOWN        = sdl.K_DOWN        // "Down" (the Down arrow key (navigation keypad))
	KEY_UP          = sdl.K_UP          // "Up" (the Up arrow key (navigation keypad))

	KEY_NUMLOCKCLEAR = sdl.K_NUMLOCKCLEAR // "Numlock" (the Num Lock key (PC) / the Clear key (Mac))
	KEY_KP_DIVIDE    = sdl.K_KP_DIVIDE    // "Keypad /" (the / key (numeric keypad))
	KEY_KP_MULTIPLY  = sdl.K_KP_MULTIPLY  // "Keypad *" (the * key (numeric keypad))
	KEY_KP_MINUS     = sdl.K_KP_MINUS     // "Keypad -" (the - key (numeric keypad))
	KEY_KP_PLUS      = sdl.K_KP_PLUS      // "Keypad +" (the + key (numeric keypad))
	KEY_KP_ENTER     = sdl.K_KP_ENTER     // "Keypad Enter" (the Enter key (numeric keypad))
	KEY_KP_1         = sdl.K_KP_1         // "Keypad 1" (the 1 key (numeric keypad))
	KEY_KP_2         = sdl.K_KP_2         // "Keypad 2" (the 2 key (numeric keypad))
	KEY_KP_3         = sdl.K_KP_3         // "Keypad 3" (the 3 key (numeric keypad))
	KEY_KP_4         = sdl.K_KP_4         // "Keypad 4" (the 4 key (numeric keypad))
	KEY_KP_5         = sdl.K_KP_5         // "Keypad 5" (the 5 key (numeric keypad))
	KEY_KP_6         = sdl.K_KP_6         // "Keypad 6" (the 6 key (numeric keypad))
	KEY_KP_7         = sdl.K_KP_7         // "Keypad 7" (the 7 key (numeric keypad))
	KEY_KP_8         = sdl.K_KP_8         // "Keypad 8" (the 8 key (numeric keypad))
	KEY_KP_9         = sdl.K_KP_9         // "Keypad 9" (the 9 key (numeric keypad))
	KEY_KP_0         = sdl.K_KP_0         // "Keypad 0" (the 0 key (numeric keypad))
	KEY_KP_PERIOD    = sdl.K_KP_PERIOD    // "Keypad ." (the . key (numeric keypad))

	KEY_APPLICATION    = sdl.K_APPLICATION    // "Application" (the Application / Compose / Context Menu (Windows) key)
	KEY_POWER          = sdl.K_POWER          // "Power" (The USB document says this is a status flag, not a physical key - but some Mac keyboards do have a power key.)
	KEY_KP_EQUALS      = sdl.K_KP_EQUALS      // "Keypad =" (the = key (numeric keypad))
	KEY_F13            = sdl.K_F13            // "F13"
	KEY_F14            = sdl.K_F14            // "F14"
	KEY_F15            = sdl.K_F15            // "F15"
	KEY_F16            = sdl.K_F16            // "F16"
	KEY_F17            = sdl.K_F17            // "F17"
	KEY_F18            = sdl.K_F18            // "F18"
	KEY_F19            = sdl.K_F19            // "F19"
	KEY_F20            = sdl.K_F20            // "F20"
	KEY_F21            = sdl.K_F21            // "F21"
	KEY_F22            = sdl.K_F22            // "F22"
	KEY_F23            = sdl.K_F23            // "F23"
	KEY_F24            = sdl.K_F24            // "F24"
	KEY_EXECUTE        = sdl.K_EXECUTE        // "Execute"
	KEY_HELP           = sdl.K_HELP           // "Help"
	KEY_MENU           = sdl.K_MENU           // "Menu"
	KEY_SELECT         = sdl.K_SELECT         // "Select"
	KEY_STOP           = sdl.K_STOP           // "Stop"
	KEY_AGAIN          = sdl.K_AGAIN          // "Again" (the Again key (Redo))
	KEY_UNDO           = sdl.K_UNDO           // "Undo"
	KEY_CUT            = sdl.K_CUT            // "Cut"
	KEY_COPY           = sdl.K_COPY           // "Copy"
	KEY_PASTE          = sdl.K_PASTE          // "Paste"
	KEY_FIND           = sdl.K_FIND           // "Find"
	KEY_MUTE           = sdl.K_MUTE           // "Mute"
	KEY_VOLUMEUP       = sdl.K_VOLUMEUP       // "VolumeUp"
	KEY_VOLUMEDOWN     = sdl.K_VOLUMEDOWN     // "VolumeDown"
	KEY_KP_COMMA       = sdl.K_KP_COMMA       // "Keypad ," (the Comma key (numeric keypad))
	KEY_KP_EQUALSAS400 = sdl.K_KP_EQUALSAS400 // "Keypad = (AS400)" (the Equals AS400 key (numeric keypad))

	KEY_ALTERASE   = sdl.K_ALTERASE   // "AltErase" (Erase-Eaze)
	KEY_SYSREQ     = sdl.K_SYSREQ     // "SysReq" (the SysReq key)
	KEY_CANCEL     = sdl.K_CANCEL     // "Cancel"
	KEY_CLEAR      = sdl.K_CLEAR      // "Clear"
	KEY_PRIOR      = sdl.K_PRIOR      // "Prior"
	KEY_RETURN2    = sdl.K_RETURN2    // "Return"
	KEY_SEPARATOR  = sdl.K_SEPARATOR  // "Separator"
	KEY_OUT        = sdl.K_OUT        // "Out"
	KEY_OPER       = sdl.K_OPER       // "Oper"
	KEY_CLEARAGAIN = sdl.K_CLEARAGAIN // "Clear / Again"
	KEY_CRSEL      = sdl.K_CRSEL      // "CrSel"
	KEY_EXSEL      = sdl.K_EXSEL      // "ExSel"

	KEY_KP_00              = sdl.K_KP_00              // "Keypad 00" (the 00 key (numeric keypad))
	KEY_KP_000             = sdl.K_KP_000             // "Keypad 000" (the 000 key (numeric keypad))
	KEY_THOUSANDSSEPARATOR = sdl.K_THOUSANDSSEPARATOR // "ThousandsSeparator" (the Thousands Separator key)
	KEY_DECIMALSEPARATOR   = sdl.K_DECIMALSEPARATOR   // "DecimalSeparator" (the Decimal Separator key)
	KEY_CURRENCYUNIT       = sdl.K_CURRENCYUNIT       // "CurrencyUnit" (the Currency Unit key)
	KEY_CURRENCYSUBUNIT    = sdl.K_CURRENCYSUBUNIT    // "CurrencySubUnit" (the Currency Subunit key)
	KEY_KP_LEFTPAREN       = sdl.K_KP_LEFTPAREN       // "Keypad (" (the Left Parenthesis key (numeric keypad))
	KEY_KP_RIGHTPAREN      = sdl.K_KP_RIGHTPAREN      // "Keypad )" (the Right Parenthesis key (numeric keypad))
	KEY_KP_LEFTBRACE       = sdl.K_KP_LEFTBRACE       // "Keypad {" (the Left Brace key (numeric keypad))
	KEY_KP_RIGHTBRACE      = sdl.K_KP_RIGHTBRACE      // "Keypad }" (the Right Brace key (numeric keypad))
	KEY_KP_TAB             = sdl.K_KP_TAB             // "Keypad Tab" (the Tab key (numeric keypad))
	KEY_KP_BACKSPACE       = sdl.K_KP_BACKSPACE       // "Keypad Backspace" (the Backspace key (numeric keypad))
	KEY_KP_A               = sdl.K_KP_A               // "Keypad A" (the A key (numeric keypad))
	KEY_KP_B               = sdl.K_KP_B               // "Keypad B" (the B key (numeric keypad))
	KEY_KP_C               = sdl.K_KP_C               // "Keypad C" (the C key (numeric keypad))
	KEY_KP_D               = sdl.K_KP_D               // "Keypad D" (the D key (numeric keypad))
	KEY_KP_E               = sdl.K_KP_E               // "Keypad E" (the E key (numeric keypad))
	KEY_KP_F               = sdl.K_KP_F               // "Keypad F" (the F key (numeric keypad))
	KEY_KP_XOR             = sdl.K_KP_XOR             // "Keypad XOR" (the XOR key (numeric keypad))
	KEY_KP_POWER           = sdl.K_KP_POWER           // "Keypad ^" (the Power key (numeric keypad))
	KEY_KP_PERCENT         = sdl.K_KP_PERCENT         // "Keypad %" (the Percent key (numeric keypad))
	KEY_KP_LESS            = sdl.K_KP_LESS            // "Keypad <" (the Less key (numeric keypad))
	KEY_KP_GREATER         = sdl.K_KP_GREATER         // "Keypad >" (the Greater key (numeric keypad))
	KEY_KP_AMPERSAND       = sdl.K_KP_AMPERSAND       // "Keypad &" (the & key (numeric keypad))
	KEY_KP_DBLAMPERSAND    = sdl.K_KP_DBLAMPERSAND    // "Keypad &&" (the && key (numeric keypad))
	KEY_KP_VERTICALBAR     = sdl.K_KP_VERTICALBAR     // "Keypad |" (the | key (numeric keypad))
	KEY_KP_DBLVERTICALBAR  = sdl.K_KP_DBLVERTICALBAR  // "Keypad ||" (the || key (numeric keypad))
	KEY_KP_COLON           = sdl.K_KP_COLON           // "Keypad :" (the : key (numeric keypad))
	KEY_KP_HASH            = sdl.K_KP_HASH            // "Keypad #" (the # key (numeric keypad))
	KEY_KP_SPACE           = sdl.K_KP_SPACE           // "Keypad Space" (the Space key (numeric keypad))
	KEY_KP_AT              = sdl.K_KP_AT              // "Keypad @" (the @ key (numeric keypad))
	KEY_KP_EXCLAM          = sdl.K_KP_EXCLAM          // "Keypad !" (the ! key (numeric keypad))
	KEY_KP_MEMSTORE        = sdl.K_KP_MEMSTORE        // "Keypad MemStore" (the Mem Store key (numeric keypad))
	KEY_KP_MEMRECALL       = sdl.K_KP_MEMRECALL       // "Keypad MemRecall" (the Mem Recall key (numeric keypad))
	KEY_KP_MEMCLEAR        = sdl.K_KP_MEMCLEAR        // "Keypad MemClear" (the Mem Clear key (numeric keypad))
	KEY_KP_MEMADD          = sdl.K_KP_MEMADD          // "Keypad MemAdd" (the Mem Add key (numeric keypad))
	KEY_KP_MEMSUBTRACT     = sdl.K_KP_MEMSUBTRACT     // "Keypad MemSubtract" (the Mem Subtract key (numeric keypad))
	KEY_KP_MEMMULTIPLY     = sdl.K_KP_MEMMULTIPLY     // "Keypad MemMultiply" (the Mem Multiply key (numeric keypad))
	KEY_KP_MEMDIVIDE       = sdl.K_KP_MEMDIVIDE       // "Keypad MemDivide" (the Mem Divide key (numeric keypad))
	KEY_KP_PLUSMINUS       = sdl.K_KP_PLUSMINUS       // "Keypad +/-" (the +/- key (numeric keypad))
	KEY_KP_CLEAR           = sdl.K_KP_CLEAR           // "Keypad Clear" (the Clear key (numeric keypad))
	KEY_KP_CLEARENTRY      = sdl.K_KP_CLEARENTRY      // "Keypad ClearEntry" (the Clear Entry key (numeric keypad))
	KEY_KP_BINARY          = sdl.K_KP_BINARY          // "Keypad Binary" (the Binary key (numeric keypad))
	KEY_KP_OCTAL           = sdl.K_KP_OCTAL           // "Keypad Octal" (the Octal key (numeric keypad))
	KEY_KP_DECIMAL         = sdl.K_KP_DECIMAL         // "Keypad Decimal" (the Decimal key (numeric keypad))
	KEY_KP_HEXADECIMAL     = sdl.K_KP_HEXADECIMAL     // "Keypad Hexadecimal" (the Hexadecimal key (numeric keypad))

	KEY_LCTRL  = sdl.K_LCTRL  // "Left Ctrl"
	KEY_LSHIFT = sdl.K_LSHIFT // "Left Shift"
	KEY_LALT   = sdl.K_LALT   // "Left Alt" (alt, option)
	KEY_LGUI   = sdl.K_LGUI   // "Left GUI" (windows, command (apple), meta)
	KEY_RCTRL  = sdl.K_RCTRL  // "Right Ctrl"
	KEY_RSHIFT = sdl.K_RSHIFT // "Right Shift"
	KEY_RALT   = sdl.K_RALT   // "Right Alt" (alt, option)
	KEY_RGUI   = sdl.K_RGUI   // "Right GUI" (windows, command (apple), meta)

	KEY_MODE = sdl.K_MODE // "ModeSwitch" (I'm not sure if this is really not covered by any of the above, but since there's a special KMOD_MODE for it I'm adding it here)

	KEY_AUDIONEXT    = sdl.K_AUDIONEXT    // "AudioNext" (the Next Track media key)
	KEY_AUDIOPREV    = sdl.K_AUDIOPREV    // "AudioPrev" (the Previous Track media key)
	KEY_AUDIOSTOP    = sdl.K_AUDIOSTOP    // "AudioStop" (the Stop media key)
	KEY_AUDIOPLAY    = sdl.K_AUDIOPLAY    // "AudioPlay" (the Play media key)
	KEY_AUDIOMUTE    = sdl.K_AUDIOMUTE    // "AudioMute" (the Mute volume key)
	KEY_MEDIASELECT  = sdl.K_MEDIASELECT  // "MediaSelect" (the Media Select key)
	KEY_WWW          = sdl.K_WWW          // "WWW" (the WWW/World Wide Web key)
	KEY_MAIL         = sdl.K_MAIL         // "Mail" (the Mail/eMail key)
	KEY_CALCULATOR   = sdl.K_CALCULATOR   // "Calculator" (the Calculator key)
	KEY_COMPUTER     = sdl.K_COMPUTER     // "Computer" (the My Computer key)
	KEY_AC_SEARCH    = sdl.K_AC_SEARCH    // "AC Search" (the Search key (application control keypad))
	KEY_AC_HOME      = sdl.K_AC_HOME      // "AC Home" (the Home key (application control keypad))
	KEY_AC_BACK      = sdl.K_AC_BACK      // "AC Back" (the Back key (application control keypad))
	KEY_AC_FORWARD   = sdl.K_AC_FORWARD   // "AC Forward" (the Forward key (application control keypad))
	KEY_AC_STOP      = sdl.K_AC_STOP      // "AC Stop" (the Stop key (application control keypad))
	KEY_AC_REFRESH   = sdl.K_AC_REFRESH   // "AC Refresh" (the Refresh key (application control keypad))
	KEY_AC_BOOKMARKS = sdl.K_AC_BOOKMARKS // "AC Bookmarks" (the Bookmarks key (application control keypad))

	KEY_BRIGHTNESSDOWN = sdl.K_BRIGHTNESSDOWN // "BrightnessDown" (the Brightness Down key)
	KEY_BRIGHTNESSUP   = sdl.K_BRIGHTNESSUP   // "BrightnessUp" (the Brightness Up key)
	KEY_DISPLAYSWITCH  = sdl.K_DISPLAYSWITCH  // "DisplaySwitch" (display mirroring/dual display switch, video mode switch)
	KEY_KBDILLUMTOGGLE = sdl.K_KBDILLUMTOGGLE // "KBDIllumToggle" (the Keyboard Illumination Toggle key)
	KEY_KBDILLUMDOWN   = sdl.K_KBDILLUMDOWN   // "KBDIllumDown" (the Keyboard Illumination Down key)
	KEY_KBDILLUMUP     = sdl.K_KBDILLUMUP     // "KBDIllumUp" (the Keyboard Illumination Up key)
	KEY_EJECT          = sdl.K_EJECT          // "Eject" (the Eject key)
	KEY_SLEEP          = sdl.K_SLEEP          // "Sleep" (the Sleep key)
)
