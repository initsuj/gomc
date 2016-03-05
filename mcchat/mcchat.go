package mcchat

type color string

func (c color) Color() color {
	return c
}

type Colorer interface {
	Color() color
}

const (
	Black      color = "§0"
	DarkBlue   color = "§1"
	DarkGreen  color = "§2"
	DarkAqua   color = "§3"
	DarkRed    color = "§4"
	DarkPurple color = "§5"
	Gold       color = "§6"
	Gray       color = "§7"
	DarkGray   color = "§8"
	Blue       color = "§9"
	Green      color = "§a"
	Aqua       color = "§b"
	Red        color = "§c"
	Purple     color = "§d"
	Yellow     color = "§e"
	White      color = "§f"
)

type format string

func (f format) Format() format {
	return f
}

type Formatter interface {
	Format() format
}

const (
	Obfuscated    format = "§k"
	Bold          format = "§l"
	Strikethrough format = "§m"
	Underline     format = "§n"
	Italic        format = "§o"
	Reset         format = "§r"
)
