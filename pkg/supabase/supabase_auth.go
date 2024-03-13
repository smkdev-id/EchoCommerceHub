package supabase

import (
	"context"
	"fmt"

	supa "github.com/nedpals/supabase-go"
	"smkdev-id/promotion_tracking_dashboard/internal/configs"
	"github.com/spf13/viper"
)

type SupabaseClient interface {
	SupabaseSignUp()
	SupabaseSignIn()
	SupabaseInsertData()
	SupabaseSelectData()
	SupabaseUpdateData()
	SupabaseDeleteData()
}

// type SupabaseClientImpl struct {
// 	supa *supabase.Client
// }

// func NewSupabaseClient(supa *supabase.Client) SupabaseClient {
// 	return &SupabaseClientImpl{
// 		supa: supa,
// 	}
// }

// func (supa *SupabaseClientImpl) SupabaseSignUp() {
func SupabaseSignUp() {
	configs.LoadViperEnv()

	supabaseUrl := viper.GetString("SUPABASE.URL")
	supabaseKey := viper.GetString("SUPABASE.KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	ctx := context.Background()
	user, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
		Email:    "samuel@eudeka.id",
		Password: "testestest",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

// func (s *SupabaseClientImpl) SupabaseSignIn() {
func SupabaseSignIn() {
	configs.LoadViperEnv()

	supabaseUrl := viper.GetString("SUPABASE.URL")
	supabaseKey := viper.GetString("SUPABASE.KEY")
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	ctx := context.Background()
	user, err := supabase.Auth.SignIn(ctx, supa.UserCredentials{
		Email:    "samuel@eudeka.id",
		Password: "testestest",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(user)
}

// func (s *SupabaseClientImpl) SupabaseInsertData() {

// }

// func (s *SupabaseClientImpl) SupabaseSelectData() {

// }

// func (s *SupabaseClientImpl) SupabaseUpdateData() {

// }

// func (s *SupabaseClientImpl) SupabaseDeleteData() {

// }

// func SupabaseAuth() {
// 	configs.LoadViperEnv()

// 	supabaseUrl := viper.GetString("SUPABASE.URL")
// 	supabaseKey := viper.GetString("SUPABASE.KEY")
// 	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

// 	ctx := context.Background()
// 	user, err := supabase.Auth.SignUp(ctx, supa.UserCredentials{
// 		Email:    "samuel@eudeka.id",
// 		Password: "testestest",
// 	})
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(user)
// }
