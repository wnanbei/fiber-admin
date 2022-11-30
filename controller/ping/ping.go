package ping

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wnanbei/fiber-admin/internal/validate"
)

// PingParams 示例参数
type PingParams struct {
	Number string `json:"number" validate:"required"` // 示例字段
}

// Ping 请求响应
//
//	@Summary		请求响应
//	@Description	请求响应
//	@Tags			test
//	@Accept			json
//	@Produce		json
//	@Param			id	query		PingParams	true	"params"
//	@Success		200	{object}	res.Res
//	@Failure		404	{object}	res.Res
//	@Router			/ping [GET]
func Ping(ctx *fiber.Ctx) error {
	var params PingParams
	if err := ctx.QueryParser(&params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}
	if err := validate.Do(params); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err)
	}

	return ctx.Status(fiber.StatusOK).SendString(params.Number)
}
