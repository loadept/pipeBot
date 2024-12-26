package throwable

import "errors"

var MisusedCommand = errors.New("The way to use this command is as follows: `|role+ @user role`.")
var NoMentions = errors.New("You need to mention a user to be able to execute the command.")
var MultipleMentions = errors.New("You can only mention one user to execute the command.")
