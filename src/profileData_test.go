package src

import (
	"log"
	"reflect"
	"testing"

	"github.com/InstagramExtractor/utils"
)

func TestGetInstaData(t *testing.T) {

	h := utils.DownloadHtml("https://www.instagram.com/google/")
	log.Println(h)

	type args struct {
		content string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "Test 1",
			args: args{
				content: h,
			},
			want: User{
				Id:           "389801252",
				Url:          "https://www.instagram.com/google/",
				Username:     "google",
				Description:  "Google unfiltered\u2014sometimes with filters.",
				ProfileImage: "https://instagram.fbcn3-1.fna.fbcdn.net/v/t51.2885-19/s150x150/119515245_239175997499686_2853342285794408974_n.jpg?_nc_ht=instagram.fbcn3-1.fna.fbcdn.net&_nc_ohc=DTvNe6hgkQQAX_ncg46&oh=c675695f0e0a31152023f3a3f6b2dd31&oe=5FC848DB",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetInstaData(tt.args.content); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInstaData() = %v, want %v", got, tt.want)
			}
		})
	}
}
