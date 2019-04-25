package file

import "os"

type Options struct {
	UID         int
	GID         int
	Flags       int
	Contents    string
	Permissions os.FileMode
}

type Option func(*Options)

func UID(userID int) Option {
	return func(args *Options) {
		args.UID = userID
	}
}

func GID(groupID int) Option {
	return func(args *Options) {
		args.GID = groupID
	}
}

func Contents(c string) Option {
	return func(args *Options) {
		args.Contents = c
	}
}

func Permissions(perms os.FileMode) Option {
	return func(args *Options) {
		args.Permissions = perms
	}
}
