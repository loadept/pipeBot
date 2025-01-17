package throwable

import "errors"

var (
	MisusedCommand               = errors.New("The way to use this command is as follows: `|role+ @user role`.")
	NoMentions                   = errors.New("You need to mention a user to be able to execute the command.")
	NoMentionsRole               = errors.New("You need to mention a role to be able to execute the command.")
	MultipleMentions             = errors.New("You can only mention one user to execute the command.")
	MultipleMentionsRole         = errors.New("You can only mention one role to execute the command.")
	WithoutSufficientPermissions = errors.New("You do not have permission to perform this action.")
	RoleDoesNotExists            = errors.New("The role you tried to assign does not exist.")
	SomethingWentWrong           = errors.New("Something went wrong. This action could not be performed.")
	SomethingWentWrongMember     = errors.New("Something went wrong. Could not get member information.")
	SomethingWentWrongRole       = errors.New("Something went wrong. Failed to get role information.")
)
