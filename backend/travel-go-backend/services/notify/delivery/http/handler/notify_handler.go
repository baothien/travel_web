package handler

import (
	"encoding/json"
	"fmt"
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/tidwall/gjson"
	"gitlab.com/virtual-travel/travel-go-backend/infrastructure/logger"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/delivery/http/model/request"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/entity"
	"gitlab.com/virtual-travel/travel-go-backend/services/notify/domain/notify/usecase"
	"gitlab.com/virtual-travel/travel-go-backend/utils/apiutil/api_response"
	"gitlab.com/virtual-travel/travel-go-backend/utils/i18nutil"
	"gitlab.com/virtual-travel/travel-go-backend/utils/randutil"
	"go.uber.org/zap"
	"log"
	"net/http"
)

type NotifyHandler struct {
	NotifyUseCase usecase.NotifyUseCase
	Hub           *entity.Hub
}

func NewNotifyHandler(useCase usecase.NotifyUseCase, hub entity.Hub) NotifyHandler {
	return NotifyHandler{
		NotifyUseCase: useCase,
		Hub:           &hub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// serveWs handles websocket requests from the peer.
func serveWs(hub *entity.Hub, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	fmt.Println(id)

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := &entity.Client{ID: id, Hub: hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()
}

func (n NotifyHandler) SocketNotify(c *gin.Context) {

	id := c.Param("id")

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &entity.Client{ID: id, Hub: n.Hub, Conn: conn, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.WritePump()
	go client.ReadPump()

}

// @Summary Api push thông  báo
// @Tags notify
// @description Api push thông  báo
// @Accept  json
// @Produce  json
// @Param data body entity.Notify true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /notify/push [post]
func (n NotifyHandler) PushNotifySocket(c *gin.Context) {
	//
	notify := entity.Notify{}
	if err := c.Bind(&notify); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	notiStr, _ := json.Marshal(notify)

	n.Hub.PushMessage(notify.ToUserID, notiStr)

}

// @Summary Api tạo thông báo
// @Tags notify
// @description Api tạo thông báo
// @Accept  json
// @Produce  json
// @Param data body entity.Notify true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /notify/create [post]
func (n NotifyHandler) CreateNotify(c *gin.Context) {
	body := entity.Notify{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.ID = randutil.UUIDRand()

	err := n.NotifyUseCase.CreateNotify(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	// get info user
	client := resty.New()
	respUser, errUser := client.R().
		SetAuthToken(c.Value("JWT_TOKEN").(string)).
		Get(fmt.Sprintf("https://travel-api.huytx.com/stag/user-service/user/detail/%s", body.FromUserID))

	if errUser != nil {
		logger.Error("Error get info user", zap.Error(errUser))
	}

	esultGjson := gjson.Get(string(respUser.Body()), "error_info")
	if esultGjson.String() != "" {
		logger.Error("Error push notify", zap.String("err", esultGjson.String()))

	}

	resultGjson := gjson.Get(string(respUser.Body()), "data")
	if resultGjson.String() != "" {

	}

	json.Unmarshal([]byte(resultGjson.String()), &body.FromUser)

	notiStr, _ := json.Marshal(body)

	n.Hub.PushMessage(body.ToUserID, notiStr)

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Tạo thông báo thành công"))
}

// @Summary Api danh sách thông báo
// @Tags notify
// @description Api danh sách thông báo
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param limit query string false "Limit record for getting"
// @Param page query string false "Current page"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /notify/list [get]
func (n NotifyHandler) GetNotifyPaging(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	body := request.ListNotifyReq{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.UserID = userId

	notifies, err := n.NotifyUseCase.GetNotifyPaging(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(notifies, ""))
}

// @Summary Api update thông báo đã đọc
// @Tags notify
// @description Api update thông báo đã đọc
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Param id path string true "id"
// @Param data body request.UpdateIsRead true "request"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /notify/is-read/{id} [patch]
func (n NotifyHandler) UpdateIsReadNotify(c *gin.Context) {
	id := c.Param("id")

	body := request.UpdateIsRead{}
	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, api_response.SimpleErrorResponse(http.StatusBadRequest, err.Error()))
		return
	}

	body.ID = id

	err := n.NotifyUseCase.UpdateIsRead(c, body)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}

	c.JSON(http.StatusOK, api_response.NewSuccessResponse(body, "Cập nhật thành công"))
}

// @Summary Api lấy số thông báo chưa xem
// @Tags notify
// @description Api lấy số thông báo chưa xem
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer token"
// @Success 200 {object} api_response.Response
// @Failure 400 {object} api_response.Response
// @Failure 500 {object} api_response.Response
// @Router /notify/count [get]
func (n NotifyHandler) GetCountNotify(c *gin.Context) {
	claims := jwt.ExtractClaims(c)
	userId := claims["UserId"].(string)

	notifies, err := n.NotifyUseCase.CountNotifyRead(c, userId)
	if err != nil {
		message, _ := i18nutil.GetI18nMessage(err.Error(), c)
		c.JSON(http.StatusInternalServerError, api_response.SimpleErrorResponse(http.StatusInternalServerError, message))
		return
	}
	c.JSON(http.StatusOK, api_response.NewSuccessResponse(notifies, ""))
}
