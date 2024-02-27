// Argo is a simple package designed to parse strings
// into command-line arguments following the
// POSIX Shell Command Language.
package argo

// Parse parses string s compliant with POSIX standard.
// Returns arguments as string array with size 10, and
// incomplete as boolean indicating whether or not string
// s is valid, or needs to be completed. If incomplete
// is true, arguments will be empty.
func Parse(s string) (arguments []string, incomplete bool) {
	var args = make([]string, 0, 10)

	if !Terminates(s) {
		return args, true
	}

	var escape bool
	var quote string
	var arg string
	for _, r := range s {
		if r == '\\' {
			escape = true
			continue
		}

		if escape {
			switch r {
			case 'a':
				arg += "\a"
			case 'b':
				arg += "\b"
			case 'f':
				arg += "\f"
			case 'n':
				arg += "\n"
			case 'r':
				arg += "\r"
			case 't':
				arg += "\t"
			case 'v':
				arg += "\v"

			case '|', '&', ';', '<', '>', '(', ')', '$', '`', '\\', '"', '\'', ' ', '\t', '\n':
				arg += string(r)

			default:
				arg += "\\" + string(r)
			}
		} else {
			switch r {
			case ' ', '\t':
				if quote == "\"" {
					switch r {
					case '|', '&', ';', '<', '>', '(', ')', '$', '`', '\\', '\'', ' ', '\t', '\n':
						arg += string(r)
					}
				} else if quote == "'" {
					switch r {
					case '|', '&', ';', '<', '>', '(', ')', '$', '`', '\\', '"', ' ', '\t', '\n':
						arg += string(r)
					}
				} else if arg != "" {
					args = append(args, arg)
					arg = ""
				}
			case '"':
				if quote == "" {
					quote = "\""
				} else if quote == "\"" {
					quote = ""
				}
			case '\'':
				if quote == "" {
					quote = "'"
				} else if quote == "'" {
					quote = ""
				}
			default:
				arg += string(r)
			}
		}

		escape = false
	}
	if arg != "" {
		args = append(args, arg)
	}

	return args, false
}

// Terminates takes in string s and checks to see if
// quotes or double quotes terminate. Returns a boolean.
func Terminates(s string) bool {
	var terminating string
	for _, r := range s {
		switch r {
		case '"':
			if terminating == "" {
				terminating = "\""
			} else if terminating == "\"" {
				terminating = ""
			}
		case '\'':
			if terminating == "" {
				terminating = "'"
			} else if terminating == "'" {
				terminating = ""
			}
		}
	}
	return terminating == ""
}
