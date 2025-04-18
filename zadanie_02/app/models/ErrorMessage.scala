package models

import play.api.libs.json.{Format, Json}

case class ErrorMessage(error: String)

object ErrorMessage {
  implicit val errorMessageFormat: Format[ErrorMessage] = Json.format[ErrorMessage]
}
