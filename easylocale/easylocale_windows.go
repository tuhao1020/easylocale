package easylocale

const (
	nameMaxLength = 85
	nullptr       = 0
	dllName       = "kernel32"
)

var procNames = []string{"GetUserDefaultLocaleName", "GetSystemDefaultLocaleName"}

func getWindowsLocaleFromProc(procName string) (string, error) {
	dll := windows.MustLoadDLL(dllName)
	proc := dll.MustFindProc(procName)

	buffer := make([]uint16, nameMaxLength)
	r, _, err := proc.Call(uintptr(unsafe.Pointer(&buffer[0])), uintptr(nameMaxLength))
	if r == nullptr {
		return "", fmt.Errorf("invoke %s: %v", procName, err)
	}

	return windows.UTF16ToString(buffer), nil
}

// CurrentLocale get locale through system call
func CurrentLocale() (string, error) {
	var locale string
	var err error

	for _, proc := range procNames {
		locale, err = getWindowsLocaleFromProc(proc)
		if err != nil {
			continue
		}
	}

	return locale, err
}
