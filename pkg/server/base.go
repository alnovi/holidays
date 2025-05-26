package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HttpController interface {
	ApplyHTTP(group *echo.Group)
}

func StatusText(code int) string {
	mapStatus := map[int]string{
		http.StatusContinue:                      "Продолжайте",
		http.StatusSwitchingProtocols:            "Переключение протоколов",
		http.StatusProcessing:                    "Идёт обработка",
		http.StatusEarlyHints:                    "Ранняя метаинформация",
		http.StatusOK:                            "Успешно",
		http.StatusCreated:                       "Создано",
		http.StatusAccepted:                      "Принято",
		http.StatusNonAuthoritativeInfo:          "Информация не авторитетна",
		http.StatusNoContent:                     "Нет содержимого",
		http.StatusResetContent:                  "Сбросить содержимое",
		http.StatusPartialContent:                "Частичное содержимое",
		http.StatusMultiStatus:                   "Многостатусный",
		http.StatusAlreadyReported:               "Уже сообщалось",
		http.StatusIMUsed:                        "Использовано IM",
		http.StatusMultipleChoices:               "Множество выборов",
		http.StatusMovedPermanently:              "Перемещено навсегда",
		http.StatusFound:                         "Перемещено временно",
		http.StatusSeeOther:                      "Смотреть другое",
		http.StatusNotModified:                   "Не изменялось",
		http.StatusUseProxy:                      "Использовать прокси",
		http.StatusTemporaryRedirect:             "Временное перенаправление",
		http.StatusPermanentRedirect:             "Постоянное перенаправление",
		http.StatusBadRequest:                    "Некорректный запрос",
		http.StatusUnauthorized:                  "Не авторизован",
		http.StatusPaymentRequired:               "Необходима оплата",
		http.StatusForbidden:                     "Доступ запрещен",
		http.StatusNotFound:                      "Не найдено",
		http.StatusMethodNotAllowed:              "Метод не поддерживается",
		http.StatusNotAcceptable:                 "Неприемлемо",
		http.StatusProxyAuthRequired:             "Необходима аутентификация прокси",
		http.StatusRequestTimeout:                "Истекло время ожидания",
		http.StatusConflict:                      "Конфликт",
		http.StatusGone:                          "Удалён",
		http.StatusLengthRequired:                "Необходима длина",
		http.StatusPreconditionFailed:            "Условие ложно",
		http.StatusRequestEntityTooLarge:         "Полезная нагрузка слишком велика",
		http.StatusRequestURITooLong:             "URI слишком длинный",
		http.StatusUnsupportedMediaType:          "Неподдерживаемый тип данных",
		http.StatusRequestedRangeNotSatisfiable:  "Диапазон не достижим",
		http.StatusExpectationFailed:             "Ожидание не удалось",
		http.StatusTeapot:                        "Я чайник",
		http.StatusMisdirectedRequest:            "Неверно направленный запрос",
		http.StatusUnprocessableEntity:           "Ошибка ввода данных",
		http.StatusLocked:                        "Доступ времмено заблокирован",
		http.StatusFailedDependency:              "Невыполненная зависимость",
		http.StatusTooEarly:                      "Слишком рано",
		http.StatusUpgradeRequired:               "Необходимо обновление",
		http.StatusPreconditionRequired:          "Необходимо предусловие",
		http.StatusTooManyRequests:               "Слишком много запросов",
		http.StatusRequestHeaderFieldsTooLarge:   "Поля заголовка запроса слишком большие",
		http.StatusUnavailableForLegalReasons:    "Недоступно по юридическим причинам",
		http.StatusInternalServerError:           "Ошибка сервера",
		http.StatusNotImplemented:                "Не реализовано",
		http.StatusBadGateway:                    "Ошибочный шлюз",
		http.StatusServiceUnavailable:            "Сервис недоступен",
		http.StatusGatewayTimeout:                "Шлюз не отвечает",
		http.StatusHTTPVersionNotSupported:       "Версия HTTP не поддерживается",
		http.StatusVariantAlsoNegotiates:         "Вариант тоже проводит согласование",
		http.StatusInsufficientStorage:           "Переполнение хранилища",
		http.StatusLoopDetected:                  "Обнаружено бесконечное перенаправление",
		http.StatusNotExtended:                   "Не расширено",
		http.StatusNetworkAuthenticationRequired: "Требуется сетевая аутентификация",
	}

	return mapStatus[code]
}
