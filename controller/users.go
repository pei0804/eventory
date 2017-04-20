package controller

import (
	"fmt"

	"github.com/goadesign/goa"
	"github.com/jinzhu/gorm"
	"github.com/tikasan/eventory/app"
	"github.com/tikasan/eventory/models"
	"github.com/tikasan/eventory/utility"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *gorm.DB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *gorm.DB) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         db,
	}
}

// ユーザーのログイン処理。
// 成功した場合は、ユーザーに登録されているtokenを返す。
// また、ログイン処理を行った端末に正規ユーザーIDを付与する。
func (c *UsersController) Login(ctx *app.LoginUsersContext) error {
	// UsersController_Login: start_implement

	// Put your logic here
	// 入力されたメールアドレスとパスワードハッシュでユーザーを見つける
	userDB := models.NewUserDB(c.db)
	user, err := userDB.UserAuth(ctx.Context, ctx.Email, ctx.PasswordHash)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	//  アカウントが正しく存在していれば該当のレコードを取得するため、tokenを取得する。
	token, err := utility.GetToken(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	userTerminalDB := models.NewUserTerminalDB(c.db)
	// tokenから、ユーザーのレコードを取得する
	ID, err := userTerminalDB.GetUserIDByToken(ctx.Context, token)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// 仮ユーザーIDに正規のユーザーIDをセットする
	userTerminal := &models.UserTerminal{}
	userTerminal.ID = ID
	userTerminal.UserID = user.ID
	// UsersController_Login: end_implement
	res := &app.Message{}
	return ctx.OK(res)
}

// 正規ユーザー作成。
// 既に登録が行われているユーザーの場合は、alreadyExistsというメッセージを返す。
// 正しく作成が行われた場合は、端末とユーザーを紐付けを行う。
func (c *UsersController) RegularCreate(ctx *app.RegularCreateUsersContext) error {
	// UsersController_RegularCreate: start_implement

	// Put your logic here
	// 正規ユーザーの登録を行う仮ユーザーを取得する
	userTerminalDB := models.NewUserTerminalDB(c.db)
	token, err := utility.GetToken(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	userTerminal, err := userTerminalDB.GetByToken(ctx.Context, token)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// メールアドレスが既に登録されているか、登録されていれば既にあることをクライアントに伝える
	userDB := models.NewUserDB(c.db)
	_, err = userDB.GetByEmail(ctx.Context, ctx.Email)
	if err == nil {
		message := "alreadyExists"
		return ctx.OK(&app.Message{&message})
	}
	// メールアドレスとユーザーを紐付ける
	newUser := &models.User{}
	newUser.ID = userTerminal.UserID
	newUser.PasswordHash = ctx.PasswordHash
	newUser.Email = ctx.Email
	err = userDB.Update(ctx.Context, newUser)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// UsersController_RegularCreate: end_implement
	message := "ok"
	return ctx.OK(&app.Message{&message})
}

// 仮ユーザー作成
// 既に登録されている端末の場合、仮ユーザーのtokenを返す。
// ユーザーを識別するtokenを発行するため、tokenの確認は行わない。
func (c *UsersController) TmpCreate(ctx *app.TmpCreateUsersContext) error {
	// UsersController_TmpCreate: start_implement

	// Put your logic here
	// 既にユーザーが存在すればtoken情報を返す
	userTerminalDB := models.NewUserTerminalDB(c.db)
	tmpUserTerminal, err := userTerminalDB.GetByIdentifier(ctx.Context, ctx.Identifier)
	if err == nil {
		return ctx.OK(&app.Token{&tmpUserTerminal.Token})
	}
	// tokenに被りがないかチェックする
	var token string
	for {
		token = utility.CreateToken(ctx.Identifier)
		_, err = userTerminalDB.GetUserIDByToken(ctx, token)
		// tokenに紐づくユーザーがいなければ終了する
		if err == nil {
			break
		}
	}
	// ユーザー作成
	t := c.db.Begin()
	userDB := models.NewUserDB(t)
	user := &models.User{}
	err = userDB.Add(ctx.Context, user)
	if err != nil {
		t.Rollback()
		return fmt.Errorf("%v", err)
	}
	userTerminalDB = models.NewUserTerminalDB(t)
	userTerminal := &models.UserTerminal{}
	userTerminal.UserID = user.ID
	userTerminal.Token = token
	userTerminal.Identifier = ctx.Identifier
	userTerminal.Platform = ctx.Platform
	userTerminal.ClientVersion = ctx.ClientVersion
	err = userTerminalDB.Add(ctx.Context, userTerminal)
	if err != nil {
		t.Rollback()
		return fmt.Errorf("%v", err)
	}
	t.Commit()
	// UsersController_TmpCreate: end_implement
	return ctx.OK(&app.Token{&userTerminal.Token})
}

// Status runs the status action.
func (c *UsersController) Status(ctx *app.StatusUsersContext) error {
	// UsersController_Status: start_implement

	// Put your logic here
	userID, err := utility.GetUserID(ctx.Context)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	userTerminal := &models.UserTerminal{}
	userTerminal.UserID = userID
	userTerminal.ClientVersion = ctx.ClientVersion
	userTerminal.Platform = ctx.Platform
	userTerminalDB := models.NewUserTerminalDB(c.db)
	err = userTerminalDB.Update(ctx.Context, userTerminal)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	// UsersController_Status: end_implement
	return nil
}
