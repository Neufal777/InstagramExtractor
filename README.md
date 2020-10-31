# Instagram Data Extractor 

InstagramExtractor is a Golang library to get data from instagram profiles.

## Installation

Import the package running the command down below

```bash
go get github.com/Neufal777/InstagramExtractor
```

## Usage

```go
import insta "github.com/Neufal777/InstagramExtractor/src"


func main() {
 
/*
Run the function(s) you need, in the readme you'll find all
the available functions
*/

  insta.DownloadPage("https://www.instagram.com/neufal79/") //takes profile url as a parameter
}
```

## Output
### _DownloadPage(ProfileUrl [string])_
This function returns a struct called `User` with the data
```bash
{
    Id: 389801252 
    Url: https://www.instagram.com/neufal79/ 
    Username: neufal79 
    Description: Description of the profile 
    ProfileImage: https://instagram.fbcn3-1.fna.fbcdn.net/v/...
}
```

_More functions coming soon.._