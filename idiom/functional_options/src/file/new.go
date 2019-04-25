package file

import "os"

func New(filepath string, setters ...Option) (*os.File, error) {
	// Default Options
	args := &Options{
		UID:         os.Getuid(),
		GID:         os.Getgid(),
		Contents:    "",
		Permissions: 0666,
		Flags:       os.O_CREATE | os.O_EXCL | os.O_WRONLY,
	}

	for _, setter := range setters {
		setter(args)
	}

	f, err := os.OpenFile(filepath, args.Flags, args.Permissions)
	if err != nil {
		return nil, err
	} else {
		defer f.Close()
	}

	if _, err := f.WriteString(args.Contents); err != nil {
		return nil, err
	}

	if err := f.Chown(args.UID, args.GID); err != nil {
		return nil, err
	}

	return f, nil
}
