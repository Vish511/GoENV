# GoENV
Simple and configurable environment config file that follows DRY principle.

Usage:
Add all the variables to be loaded from system environment to Configuration struct and provide the name of the environment variable in 'env' tag. For a fallback, provide a default value in the 'default' tag.

Example :
To add a variable to be loaded from environment variables for example, "PORT" add the variable to the 'env' tag. It will load from the environment variable that is specified in the 'env' tag.
Specify a default value to be loaded in the 'default' tag.

Typically in a production environment, port and all other application critical variables are stored in the environment variables to prevent accidental expose to developers.
So the default values are mentioned in 'default' tag. Developers will be able to see only this.
