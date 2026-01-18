package swagger

// NOTE:
// - This package is ONLY for swaggo/swag annotations.
// - These functions are NOT used at runtime.
// - We keep them separate from controllers to avoid coupling doc comments to business code.

// ----------------------------
// /api (Management API)
// ----------------------------

// @Tags system
// @Summary Get service status (public)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/status [get]
func ApiStatus() {}

// @Tags system
// @Summary List models for dashboard (requires user auth)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/models [get]
func ApiModels() {}

// @Tags system
// @Summary Get notice content (public)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/notice [get]
func ApiNotice() {}

// @Tags system
// @Summary Get about content (public)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/about [get]
func ApiAbout() {}

// @Tags system
// @Summary Get home page content (public)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/home_page_content [get]
func ApiHomePageContent() {}

// @Tags auth
// @Summary Send email verification code
// @Produce json
// @Param email query string true "Email"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/verification [get]
func ApiVerification() {}

// @Tags auth
// @Summary Send password reset email
// @Produce json
// @Param email query string true "Email"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/reset_password [get]
func ApiResetPassword() {}

// @Tags user
// @Summary Reset password
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/reset [post]
func ApiUserResetPassword() {}

// @Tags oauth
// @Summary GitHub OAuth callback
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/github [get]
func ApiOauthGitHub() {}

// @Tags oauth
// @Summary OIDC OAuth callback
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/oidc [get]
func ApiOauthOidc() {}

// @Tags oauth
// @Summary Lark OAuth callback
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/lark [get]
func ApiOauthLark() {}

// @Tags oauth
// @Summary Generate OAuth state
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/state [get]
func ApiOauthState() {}

// @Tags oauth
// @Summary WeChat OAuth callback
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/wechat [get]
func ApiOauthWeChat() {}

// @Tags oauth
// @Summary Bind WeChat account
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/wechat/bind [get]
func ApiOauthWeChatBind() {}

// @Tags oauth
// @Summary Bind email
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/oauth/email/bind [get]
func ApiOauthEmailBind() {}

// @Tags admin
// @Summary Admin top-up
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/topup [post]
func ApiAdminTopup() {}

// --- /api/user ---

// @Tags user
// @Summary Register
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/register [post]
func ApiUserRegister() {}

// @Tags user
// @Summary Login
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/login [post]
func ApiUserLogin() {}

// @Tags user
// @Summary Logout
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/logout [get]
func ApiUserLogout() {}

// @Tags user
// @Summary Get user dashboard (self)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/dashboard [get]
func ApiUserDashboard() {}

// @Tags user
// @Summary Get current user (self)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/self [get]
func ApiUserSelfGet() {}

// @Tags user
// @Summary Update current user (self)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/self [put]
func ApiUserSelfUpdate() {}

// @Tags user
// @Summary Delete current user (self)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/self [delete]
func ApiUserSelfDelete() {}

// @Tags token
// @Summary Generate access token (self)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/token [get]
func ApiUserGenerateAccessToken() {}

// @Tags user
// @Summary Get affiliate code (self)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/aff [get]
func ApiUserAff() {}

// @Tags user
// @Summary Top up (self)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/topup [post]
func ApiUserTopup() {}

// @Tags model
// @Summary Get available models for current user
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/available_models [get]
func ApiUserAvailableModels() {}

// --- /api/user (admin) ---

// @Tags admin-user
// @Summary List users
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/ [get]
func ApiAdminUsersList() {}

// @Tags admin-user
// @Summary Search users
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/search [get]
func ApiAdminUsersSearch() {}

// @Tags admin-user
// @Summary Get user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/{id} [get]
func ApiAdminUserGet() {}

// @Tags admin-user
// @Summary Create user
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/ [post]
func ApiAdminUserCreate() {}

// @Tags admin-user
// @Summary Manage user
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/manage [post]
func ApiAdminUserManage() {}

// @Tags admin-user
// @Summary Update user
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/ [put]
func ApiAdminUserUpdate() {}

// @Tags admin-user
// @Summary Delete user by ID
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/user/{id} [delete]
func ApiAdminUserDelete() {}

// --- /api/option ---

// @Tags option
// @Summary Get options (root only)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/option/ [get]
func ApiOptionGet() {}

// @Tags option
// @Summary Update options (root only)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/option/ [put]
func ApiOptionUpdate() {}

// --- /api/channel ---

// @Tags channel
// @Summary List channels
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/ [get]
func ApiChannelList() {}

// @Tags channel
// @Summary Search channels
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/search [get]
func ApiChannelSearch() {}

// @Tags channel
// @Summary List all models (admin)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/models [get]
func ApiChannelModels() {}

// @Tags channel
// @Summary Get channel by ID
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/{id} [get]
func ApiChannelGet() {}

// @Tags channel
// @Summary Test channels
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/test [get]
func ApiChannelTestAll() {}

// @Tags channel
// @Summary Test channel by ID
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/test/{id} [get]
func ApiChannelTestOne() {}

// @Tags channel
// @Summary Update all channels balance
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/update_balance [get]
func ApiChannelUpdateBalanceAll() {}

// @Tags channel
// @Summary Update channel balance by ID
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/update_balance/{id} [get]
func ApiChannelUpdateBalanceOne() {}

// @Tags channel
// @Summary Add channel
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/ [post]
func ApiChannelAdd() {}

// @Tags channel
// @Summary Update channel
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/ [put]
func ApiChannelUpdate() {}

// @Tags channel
// @Summary Delete disabled channels
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/disabled [delete]
func ApiChannelDeleteDisabled() {}

// @Tags channel
// @Summary Delete channel by ID
// @Produce json
// @Param id path int true "Channel ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/channel/{id} [delete]
func ApiChannelDelete() {}

// --- /api/token ---

// @Tags token
// @Summary List tokens
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/token/ [get]
func ApiTokenList() {}

// @Tags token
// @Summary Search tokens
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/token/search [get]
func ApiTokenSearch() {}

// @Tags token
// @Summary Get token by ID
// @Produce json
// @Param id path int true "Token ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/token/{id} [get]
func ApiTokenGet() {}

// @Tags token
// @Summary Add token
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/token/ [post]
func ApiTokenAdd() {}

// @Tags token
// @Summary Update token
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/token/ [put]
func ApiTokenUpdate() {}

// @Tags token
// @Summary Delete token by ID
// @Produce json
// @Param id path int true "Token ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/token/{id} [delete]
func ApiTokenDelete() {}

// --- /api/redemption ---

// @Tags redemption
// @Summary List redemptions
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/redemption/ [get]
func ApiRedemptionList() {}

// @Tags redemption
// @Summary Search redemptions
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/redemption/search [get]
func ApiRedemptionSearch() {}

// @Tags redemption
// @Summary Get redemption by ID
// @Produce json
// @Param id path int true "Redemption ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/redemption/{id} [get]
func ApiRedemptionGet() {}

// @Tags redemption
// @Summary Add redemption
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/redemption/ [post]
func ApiRedemptionAdd() {}

// @Tags redemption
// @Summary Update redemption
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/redemption/ [put]
func ApiRedemptionUpdate() {}

// @Tags redemption
// @Summary Delete redemption by ID
// @Produce json
// @Param id path int true "Redemption ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /api/redemption/{id} [delete]
func ApiRedemptionDelete() {}

// --- /api/log ---

// @Tags log
// @Summary List logs (admin)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/ [get]
func ApiLogList() {}

// @Tags log
// @Summary Delete history logs (admin)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/ [delete]
func ApiLogDeleteHistory() {}

// @Tags log
// @Summary Get log stats (admin)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/stat [get]
func ApiLogStat() {}

// @Tags log
// @Summary Get self log stats
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/self/stat [get]
func ApiLogSelfStat() {}

// @Tags log
// @Summary Search all logs (admin)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/search [get]
func ApiLogSearchAll() {}

// @Tags log
// @Summary Get self logs
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/self [get]
func ApiLogSelf() {}

// @Tags log
// @Summary Search self logs
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/log/self/search [get]
func ApiLogSelfSearch() {}

// --- /api/group ---

// @Tags group
// @Summary List groups (admin)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /api/group/ [get]
func ApiGroupList() {}
