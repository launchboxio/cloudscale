package sdk

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type Client struct {
	http *resty.Client
}

func New(endpoint string) *Client {
	client := resty.New()
	client.SetBaseURL(endpoint)
	client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		switch response.StatusCode() {
		case http.StatusInternalServerError:
			return errors.New("an unexpected error occured")
		case http.StatusBadRequest:
			return errors.New(response.String())
		default:
			return nil
		}
	})
	return &Client{client}
}

func (c *Client) Certificates() *Certificates {
	return &Certificates{c}
}

func (c *Client) Listeners() *Listeners {
	return &Listeners{c}
}

func (c *Client) TargetGroups() *TargetGroups {
	return &TargetGroups{c}
}

//func FromFlags(input interface{}, flags *pflag.FlagSet) (interface{}, error) {
//	reflection := reflect.TypeOf(input)
//	ps := reflect.ValueOf(input)
//	for i := 0; i < reflection.NumField(); i++ {
//		field := reflection.Field(i)
//		flag, ok := field.Tag.Lookup("flag")
//		if !ok {
//			continue
//		}
//
//		switch field.Type.Kind() {
//		case reflect.String:
//			val, _ := flags.GetString(flag)
//			if val != "" {
//				ps.FieldByName(field.Name).SetString(val)
//			}
//		}
//		//fmt.Println(flag)
//	}
//	return ps, nil
//}
