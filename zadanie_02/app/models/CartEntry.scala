package models

import play.api.libs.json.{Format, Json}

case class CartEntry(productId: Long, quantity: Long)

object CartEntry {
  implicit val cartEntryFormat: Format[CartEntry] = Json.format[CartEntry]
}
