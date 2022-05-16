# sample-cord-go
Showcase of how to implement cord backend authentication using golang

## Setup requirement
If you wish to run this script, you need to set the following environment variables:
```
export CORD="{\"app_id\":\"SOME_VALUE\",\"secret\":\"SOME_VALUE\"}"
```

You can find the values on the [cord console](https://console.cord.com/applications).

## Note
Please, bear in mind that you can optionally setup the profile picture url and therefore the `UserDetails` struct would like this:
```
type UserDetails struct {
	Name  string `json:"name"`
	Email string `json:"email"`
    ProfilePictureUrl string `json:"profile_picture_url"`
}
```