package swagger

// ----------------------------
// /dashboard (OpenAI billing compatible)
// ----------------------------

// @Tags dashboard
// @Summary Get subscription (OpenAI compatible)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /dashboard/billing/subscription [get]
func DashboardSubscription() {}

// @Tags dashboard
// @Summary Get subscription (OpenAI compatible)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/dashboard/billing/subscription [get]
func V1DashboardSubscription() {}

// @Tags dashboard
// @Summary Get usage (OpenAI compatible)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /dashboard/billing/usage [get]
func DashboardUsage() {}

// @Tags dashboard
// @Summary Get usage (OpenAI compatible)
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/dashboard/billing/usage [get]
func V1DashboardUsage() {}

// ----------------------------
// /v1 (OpenAI compatible relay)
// ----------------------------

// @Tags v1
// @Summary List models
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/models [get]
func V1ModelsList() {}

// @Tags v1
// @Summary Retrieve model
// @Produce json
// @Param model path string true "Model ID"
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/models/{model} [get]
func V1ModelsRetrieve() {}

// @Tags relay
// @Summary Relay (OneAPI proxy)
// @Produce json
// @Param channelid path string true "Channel ID"
// @Param target path string true "Target path (may contain slashes)"
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/oneapi/proxy/{channelid}/{target} [post]
func V1OneapiProxy() {}

// @Tags relay
// @Summary Completions (OpenAI compatible)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/completions [post]
func V1Completions() {}

// @Tags relay
// @Summary Chat Completions (OpenAI compatible)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/chat/completions [post]
func V1ChatCompletions() {}

// @Tags relay
// @Summary Edits (OpenAI compatible)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/edits [post]
func V1Edits() {}

// @Tags relay
// @Summary Images generations (OpenAI compatible)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/images/generations [post]
func V1ImagesGenerations() {}

// @Tags relay
// @Summary Embeddings (OpenAI compatible)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/embeddings [post]
func V1Embeddings() {}

// @Tags relay
// @Summary Embeddings (legacy engines path)
// @Accept json
// @Produce json
// @Param model path string true "Model"
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/engines/{model}/embeddings [post]
func V1EnginesEmbeddings() {}

// @Tags relay
// @Summary Audio transcriptions
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/audio/transcriptions [post]
func V1AudioTranscriptions() {}

// @Tags relay
// @Summary Audio translations
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/audio/translations [post]
func V1AudioTranslations() {}

// @Tags relay
// @Summary Audio speech
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/audio/speech [post]
func V1AudioSpeech() {}

// @Tags relay
// @Summary Moderations (OpenAI compatible)
// @Accept json
// @Produce json
// @Success 200 {object} swagger.GenericResponse
// @Router /v1/moderations [post]
func V1Moderations() {}
