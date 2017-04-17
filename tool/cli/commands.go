package cli

import (
	"../client"
	"encoding/json"
	"fmt"
	"github.com/goadesign/goa"
	goaclient "github.com/goadesign/goa/client"
	uuid "github.com/goadesign/goa/uuid"
	"github.com/spf13/cobra"
	"golang.org/x/net/context"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// AppendGenreCronCommand is the command line data structure for the append genre action of cron
	AppendGenreCronCommand struct {
		PrettyPrint bool
	}

	// FixUserFollowCronCommand is the command line data structure for the fix user follow action of cron
	FixUserFollowCronCommand struct {
		PrettyPrint bool
	}

	// NewEventFetchCronCommand is the command line data structure for the new event fetch action of cron
	NewEventFetchCronCommand struct {
		PrettyPrint bool
	}

	// KeepEventsCommand is the command line data structure for the keep action of events
	KeepEventsCommand struct {
		// イベントID
		EventID int
		// キープ操作
		IsKeep      string
		PrettyPrint bool
	}

	// ListEventsCommand is the command line data structure for the list action of events
	ListEventsCommand struct {
		ID string
		// ページ(1->2->3->4)
		Page int
		// キーワード検索
		Q string
		// ソート
		Sort        string
		PrettyPrint bool
	}

	// CreateGenresCommand is the command line data structure for the create action of genres
	CreateGenresCommand struct {
		// ジャンル名
		Name        string
		PrettyPrint bool
	}

	// FollowGenresCommand is the command line data structure for the follow action of genres
	FollowGenresCommand struct {
		// ジャンルID
		GenreID     int
		PrettyPrint bool
	}

	// ListGenresCommand is the command line data structure for the list action of genres
	ListGenresCommand struct {
		// ページ(1->2->3->4)
		Page int
		// ジャンル名検索に使うキーワード
		Q string
		// ソート
		Sort        string
		PrettyPrint bool
	}

	// FollowPrefsCommand is the command line data structure for the follow action of prefs
	FollowPrefsCommand struct {
		// 都道府県ID
		PrefID      int
		PrettyPrint bool
	}

	// LoginUsersCommand is the command line data structure for the login action of users
	LoginUsersCommand struct {
		// メールアドレス
		Email string
		// パスワードハッシュ(^[a-z0-9]{64}$)
		PasswordHash string
		PrettyPrint  bool
	}

	// RegularCreateUsersCommand is the command line data structure for the regular create action of users
	RegularCreateUsersCommand struct {
		// メールアドレス
		Email string
		// パスワードハッシュ(^[a-z0-9]{64}$)
		PasswordHash string
		PrettyPrint  bool
	}

	// StatusUsersCommand is the command line data structure for the status action of users
	StatusUsersCommand struct {
		// アプリのバージョン
		ClientVersion string
		// OSとバージョン(iOS 10.2など)
		Platform    string
		PrettyPrint bool
	}

	// TmpCreateUsersCommand is the command line data structure for the tmp create action of users
	TmpCreateUsersCommand struct {
		// アプリのバージョン
		ClientVersion string
		// 識別子(android:Android_ID, ios:IDFV)
		Identifier string
		// OSとバージョン
		Platform    string
		PrettyPrint bool
	}
)

// RegisterCommands registers the resource action CLI commands.
func RegisterCommands(app *cobra.Command, c *client.Client) {
	var command, sub *cobra.Command
	command = &cobra.Command{
		Use:   "appendGenre",
		Short: `<b>イベントにジャンルを付加する<b>`,
	}
	tmp1 := new(AppendGenreCronCommand)
	sub = &cobra.Command{
		Use:   `cron ["/api/v2/cron/events/appendgenre"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp1.Run(c, args) },
	}
	tmp1.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp1.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "create",
		Short: `<b>ジャンルの新規作成</b><br>
		新しく作成するジャンル名を送信して、新規作成を行う。追加処理が完了とするとジャンルIDが返ってくるので、それを自動でフォローするようにする。<br>
		但し、ジャンルを新規作成する前に、ジャンル名を検索するフローを挟み、検索結果に出てこなかった場合に追加できるようにする。`,
	}
	tmp2 := new(CreateGenresCommand)
	sub = &cobra.Command{
		Use:   `genres ["/api/v2/genres/new"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp2.Run(c, args) },
	}
	tmp2.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp2.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "fixUserFollow",
		Short: `<b>イベントフォロー操作の確定</b><br>
		user_follow_eventsテーブルのbatch_processedをtrueに変更する`,
	}
	tmp3 := new(FixUserFollowCronCommand)
	sub = &cobra.Command{
		Use:   `cron ["/api/v2/cron/user/events/fixfollow"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp3.Run(c, args) },
	}
	tmp3.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp3.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "follow",
		Short: `follow action`,
	}
	tmp4 := new(FollowGenresCommand)
	sub = &cobra.Command{
		Use:   `genres [("/api/v2/genres/GENREID/follow"|"/api/v2/genres/GENREID/follow")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp4.Run(c, args) },
	}
	tmp4.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp4.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp5 := new(FollowPrefsCommand)
	sub = &cobra.Command{
		Use:   `prefs [("/api/v2/prefs/PREFID/follow"|"/api/v2/prefs/PREFID/follow")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp5.Run(c, args) },
	}
	tmp5.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp5.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "keep",
		Short: `<b>イベントお気に入り操作</b><br>
		isKeepがtrueだった場合はフォロー、falseの場合はアンフォローとする。<br>
		存在しないイベントへのリクエストは404エラーを返す。`,
	}
	tmp6 := new(KeepEventsCommand)
	sub = &cobra.Command{
		Use:   `events ["/api/v2/events/EVENTID/keep"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp6.Run(c, args) },
	}
	tmp6.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp6.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "list",
		Short: `list action`,
	}
	tmp7 := new(ListEventsCommand)
	sub = &cobra.Command{
		Use:   `events [("/api/v2/events/genre/ID"|"/api/v2/events/new"|"/api/v2/events/keep"|"/api/v2/events/nokeep"|"/api/v2/events/popular"|"/api/v2/events/recommend")]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp7.Run(c, args) },
	}
	tmp7.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp7.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	tmp8 := new(ListGenresCommand)
	sub = &cobra.Command{
		Use:   `genres ["/api/v2/genres"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp8.Run(c, args) },
	}
	tmp8.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp8.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "login",
		Short: `<b>ログイン認証</b><br>
		正規ユーザーのメールアドレスとパスワードのハッシュを送ることで、ユーザー認証を行う<br>
		正しくユーザー認証が完了した場合、正規ユーザーのIDを仮ユーザーIDに紐付けを行い。<br>
		ユーザーの行動を別端末で引き継ぐことが出来る。<br>`,
	}
	tmp9 := new(LoginUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/v2/users/login"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp9.Run(c, args) },
	}
	tmp9.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp9.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use:   "newEventFetch",
		Short: `<b>最新イベント情報の取得<b>`,
	}
	tmp10 := new(NewEventFetchCronCommand)
	sub = &cobra.Command{
		Use:   `cron ["/api/v2/cron/events/fetch"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp10.Run(c, args) },
	}
	tmp10.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp10.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "regularCreate",
		Short: `<b>正規ユーザーの作成</b><br>
		メールアドレスとパスワードハッシュを使って、正規ユーザーの作成を行う。<br>
		もし、既に存在するアカウントだった場合は、"alreadyExists"を返す。<br>
		正しく実行された場合は、"ok"を返す。`,
	}
	tmp11 := new(RegularCreateUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/v2/users/new"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp11.Run(c, args) },
	}
	tmp11.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp11.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "status",
		Short: `<b>ユーザーの端末情報更新</b><br>
		利用者のバージョンや端末情報を更新する。この更新処理は起動時に行われるものとする。`,
	}
	tmp12 := new(StatusUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/v2/users/status"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp12.Run(c, args) },
	}
	tmp12.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp12.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
	command = &cobra.Command{
		Use: "tmpCreate",
		Short: `<b>一時ユーザーの作成</b><br>
		初回起動時に仮ユーザーを作成する。ここで与えられるユーザーIDは、メールアドレスなどとひも付きがないため、<br>
		端末が変わるとtokenが変わるので、別端末で共有するには、正規ユーザーの登録が必要になる。`,
	}
	tmp13 := new(TmpCreateUsersCommand)
	sub = &cobra.Command{
		Use:   `users ["/api/v2/users/tmp"]`,
		Short: ``,
		RunE:  func(cmd *cobra.Command, args []string) error { return tmp13.Run(c, args) },
	}
	tmp13.RegisterFlags(sub, c)
	sub.PersistentFlags().BoolVar(&tmp13.PrettyPrint, "pp", false, "Pretty print response body")
	command.AddCommand(sub)
	app.AddCommand(command)
}

func intFlagVal(name string, parsed int) *int {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func float64FlagVal(name string, parsed float64) *float64 {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func boolFlagVal(name string, parsed bool) *bool {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func stringFlagVal(name string, parsed string) *string {
	if hasFlag(name) {
		return &parsed
	}
	return nil
}

func hasFlag(name string) bool {
	for _, arg := range os.Args[1:] {
		if strings.HasPrefix(arg, "--"+name) {
			return true
		}
	}
	return false
}

func jsonVal(val string) (*interface{}, error) {
	var t interface{}
	err := json.Unmarshal([]byte(val), &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func jsonArray(ins []string) ([]interface{}, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []interface{}
	for _, id := range ins {
		val, err := jsonVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, val)
	}
	return vals, nil
}

func timeVal(val string) (*time.Time, error) {
	t, err := time.Parse(time.RFC3339, val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func timeArray(ins []string) ([]time.Time, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []time.Time
	for _, id := range ins {
		val, err := timeVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func uuidVal(val string) (*uuid.UUID, error) {
	t, err := uuid.FromString(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func uuidArray(ins []string) ([]uuid.UUID, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []uuid.UUID
	for _, id := range ins {
		val, err := uuidVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func float64Val(val string) (*float64, error) {
	t, err := strconv.ParseFloat(val, 64)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func float64Array(ins []string) ([]float64, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []float64
	for _, id := range ins {
		val, err := float64Val(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

func boolVal(val string) (*bool, error) {
	t, err := strconv.ParseBool(val)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func boolArray(ins []string) ([]bool, error) {
	if ins == nil {
		return nil, nil
	}
	var vals []bool
	for _, id := range ins {
		val, err := boolVal(id)
		if err != nil {
			return nil, err
		}
		vals = append(vals, *val)
	}
	return vals, nil
}

// Run makes the HTTP request corresponding to the AppendGenreCronCommand command.
func (cmd *AppendGenreCronCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/cron/events/appendgenre"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.AppendGenreCron(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *AppendGenreCronCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the FixUserFollowCronCommand command.
func (cmd *FixUserFollowCronCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/cron/user/events/fixfollow"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.FixUserFollowCron(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *FixUserFollowCronCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the NewEventFetchCronCommand command.
func (cmd *NewEventFetchCronCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/cron/events/fetch"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.NewEventFetchCron(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *NewEventFetchCronCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
}

// Run makes the HTTP request corresponding to the KeepEventsCommand command.
func (cmd *KeepEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/v2/events/%v/keep", cmd.EventID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	var tmp14 *bool
	if cmd.IsKeep != "" {
		var err error
		tmp14, err = boolVal(cmd.IsKeep)
		if err != nil {
			goa.LogError(ctx, "failed to parse flag into *bool value", "flag", "--isKeep", "err", err)
			return err
		}
	}
	resp, err := c.KeepEvents(ctx, path, *tmp14)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *KeepEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var eventID int
	cc.Flags().IntVar(&cmd.EventID, "eventID", eventID, `イベントID`)
	var isKeep string
	cc.Flags().StringVar(&cmd.IsKeep, "isKeep", isKeep, `キープ操作`)
}

// Run makes the HTTP request corresponding to the ListEventsCommand command.
func (cmd *ListEventsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/events/new"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListEvents(ctx, path, intFlagVal("page", cmd.Page), stringFlagVal("q", cmd.Q), stringFlagVal("sort", cmd.Sort))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListEventsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var id string
	cc.Flags().StringVar(&cmd.ID, "id", id, ``)
	var page int
	cc.Flags().IntVar(&cmd.Page, "page", page, `ページ(1->2->3->4)`)
	var q string
	cc.Flags().StringVar(&cmd.Q, "q", q, `キーワード検索`)
	var sort string
	cc.Flags().StringVar(&cmd.Sort, "sort", sort, `ソート`)
}

// Run makes the HTTP request corresponding to the CreateGenresCommand command.
func (cmd *CreateGenresCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/genres/new"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.CreateGenres(ctx, path, cmd.Name)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *CreateGenresCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var name string
	cc.Flags().StringVar(&cmd.Name, "name", name, `ジャンル名`)
}

// Run makes the HTTP request corresponding to the FollowGenresCommand command.
func (cmd *FollowGenresCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/v2/genres/%v/follow", cmd.GenreID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.FollowGenres(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *FollowGenresCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var genreID int
	cc.Flags().IntVar(&cmd.GenreID, "genreID", genreID, `ジャンルID`)
}

// Run makes the HTTP request corresponding to the ListGenresCommand command.
func (cmd *ListGenresCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/genres"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.ListGenres(ctx, path, intFlagVal("page", cmd.Page), stringFlagVal("q", cmd.Q), stringFlagVal("sort", cmd.Sort))
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *ListGenresCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var page int
	cc.Flags().IntVar(&cmd.Page, "page", page, `ページ(1->2->3->4)`)
	var q string
	cc.Flags().StringVar(&cmd.Q, "q", q, `ジャンル名検索に使うキーワード`)
	var sort string
	cc.Flags().StringVar(&cmd.Sort, "sort", sort, `ソート`)
}

// Run makes the HTTP request corresponding to the FollowPrefsCommand command.
func (cmd *FollowPrefsCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = fmt.Sprintf("/api/v2/prefs/%v/follow", cmd.PrefID)
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.FollowPrefs(ctx, path)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *FollowPrefsCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var prefID int
	cc.Flags().IntVar(&cmd.PrefID, "prefID", prefID, `都道府県ID`)
}

// Run makes the HTTP request corresponding to the LoginUsersCommand command.
func (cmd *LoginUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/users/login"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.LoginUsers(ctx, path, cmd.Email, cmd.PasswordHash)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *LoginUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var email string
	cc.Flags().StringVar(&cmd.Email, "email", email, `メールアドレス`)
	var passwordHash string
	cc.Flags().StringVar(&cmd.PasswordHash, "password_hash", passwordHash, `パスワードハッシュ(^[a-z0-9]{64}$)`)
}

// Run makes the HTTP request corresponding to the RegularCreateUsersCommand command.
func (cmd *RegularCreateUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/users/new"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.RegularCreateUsers(ctx, path, cmd.Email, cmd.PasswordHash)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *RegularCreateUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var email string
	cc.Flags().StringVar(&cmd.Email, "email", email, `メールアドレス`)
	var passwordHash string
	cc.Flags().StringVar(&cmd.PasswordHash, "password_hash", passwordHash, `パスワードハッシュ(^[a-z0-9]{64}$)`)
}

// Run makes the HTTP request corresponding to the StatusUsersCommand command.
func (cmd *StatusUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/users/status"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.StatusUsers(ctx, path, cmd.ClientVersion, cmd.Platform)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *StatusUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var clientVersion string
	cc.Flags().StringVar(&cmd.ClientVersion, "client_version", clientVersion, `アプリのバージョン`)
	var platform string
	cc.Flags().StringVar(&cmd.Platform, "platform", platform, `OSとバージョン(iOS 10.2など)`)
}

// Run makes the HTTP request corresponding to the TmpCreateUsersCommand command.
func (cmd *TmpCreateUsersCommand) Run(c *client.Client, args []string) error {
	var path string
	if len(args) > 0 {
		path = args[0]
	} else {
		path = "/api/v2/users/tmp"
	}
	logger := goa.NewLogger(log.New(os.Stderr, "", log.LstdFlags))
	ctx := goa.WithLogger(context.Background(), logger)
	resp, err := c.TmpCreateUsers(ctx, path, cmd.ClientVersion, cmd.Identifier, cmd.Platform)
	if err != nil {
		goa.LogError(ctx, "failed", "err", err)
		return err
	}

	goaclient.HandleResponse(c.Client, resp, cmd.PrettyPrint)
	return nil
}

// RegisterFlags registers the command flags with the command line.
func (cmd *TmpCreateUsersCommand) RegisterFlags(cc *cobra.Command, c *client.Client) {
	var clientVersion string
	cc.Flags().StringVar(&cmd.ClientVersion, "client_version", clientVersion, `アプリのバージョン`)
	var identifier string
	cc.Flags().StringVar(&cmd.Identifier, "identifier", identifier, `識別子(android:Android_ID, ios:IDFV)`)
	var platform string
	cc.Flags().StringVar(&cmd.Platform, "platform", platform, `OSとバージョン`)
}
