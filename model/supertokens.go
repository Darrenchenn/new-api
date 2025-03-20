package model

import (
	"one-api/common"

	"github.com/supertokens/supertokens-golang/recipe/emailpassword"
	"github.com/supertokens/supertokens-golang/recipe/session"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty"
	"github.com/supertokens/supertokens-golang/recipe/thirdparty/tpmodels"
	"github.com/supertokens/supertokens-golang/recipe/usermetadata"
	"github.com/supertokens/supertokens-golang/supertokens"
)

type UserSuperTokens struct {
	Id   string                 `json:"id"`
	Info map[string]interface{} `json:"info"`
}

func InitSuperTokens() {
	apiBasePath := "/auth"
	websiteBasePath := "/auth"
	err := supertokens.Init(supertokens.TypeInput{
		Supertokens: &supertokens.ConnectionInfo{
			// We use try.supertokens for demo purposes.
			// At the end of the tutorial we will show you how to create
			// your own SuperTokens core instance and then update your config.
			ConnectionURI: "https://st-dev-455a69a0-02fd-11f0-b129-952e24c20aa3.aws.supertokens.io",
			APIKey:        "lO1gVvbTDboLRKATPNhRQcPZWv",
		},
		AppInfo: supertokens.AppInfo{
			AppName:         "test",
			APIDomain:       "http://localhost:3000", //note: This should be the domain of your backend API
			WebsiteDomain:   "http://localhost:4173", //todo: it should be your frontend domain
			APIBasePath:     &apiBasePath,
			WebsiteBasePath: &websiteBasePath,
		},
		RecipeList: []supertokens.Recipe{
			usermetadata.Init(nil),
			emailpassword.Init(nil),
			session.Init(nil),
			thirdparty.Init(&tpmodels.TypeInput{
				SignInAndUpFeature: tpmodels.TypeInputSignInAndUp{
					Providers: []tpmodels.ProviderInput{
						// We have provided you with development keys which you can use for testing.
						// IMPORTANT: Please replace them with your own OAuth keys for production use.
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "google",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     "1060725074195-kmeum4crr01uirfl2op9kd5acmi9jutn.apps.googleusercontent.com",
										ClientSecret: "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",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "github",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID:     "467101b197249757c71f",
										ClientSecret: "e97051221f4b6426e8fe8d51486396703012f5bd",
									},
								},
							},
						},
						{
							Config: tpmodels.ProviderConfig{
								ThirdPartyId: "apple",
								Clients: []tpmodels.ProviderClientConfig{
									{
										ClientID: "4398792-io.supertokens.example.service",
										AdditionalConfig: map[string]interface{}{
											"keyId":      "7M48Y4RYDL",
											"privateKey": "-----BEGIN PRIVATE KEY-----\nMIGTAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBHkwdwIBAQQgu8gXs+XYkqXD6Ala9Sf/iJXzhbwcoG5dMh1OonpdJUmgCgYIKoZIzj0DAQehRANCAASfrvlFbFCYqn3I2zeknYXLwtH30JuOKestDbSfZYxZNMqhF/OzdZFTV0zc5u5s3eN+oCWbnvl0hM+9IW0UlkdA\n-----END PRIVATE KEY-----",
											"teamId":     "YWQCXGJRJL",
										},
									},
								},
							},
						},
					},
				},
			}),
		},
	})

	if err != nil {
		panic(err.Error())
	}
}

func UserMetadataUpdate(userId string) error {
	return nil
}

func UserMetadataGet(userId string) (*UserSuperTokens, error) {
	metadata, err := usermetadata.GetUserMetadata(userId)
	if err != nil {
		common.SysError("failed to get user metadata: " + err.Error())
		return nil, err
	}

	//exampleValue := metadata["exampleKey"]

	return &UserSuperTokens{
		Id:   userId,
		Info: metadata,
	}, err
}

func UserMetadataDelete() {

}
