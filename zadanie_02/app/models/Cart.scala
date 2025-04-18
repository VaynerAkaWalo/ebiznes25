package models

import play.api.libs.json.{Format, Json}

case class Cart(id: Long, items: List[CartEntry])

object Cart {
  implicit val cartFormat: Format[Cart] = Json.format[Cart]

  var carts: List[Cart] = List(
    Cart(1, List(CartEntry(1, 5))),
    Cart(1, List(CartEntry(2, 5), CartEntry(1, 2)))
  )

  def all(): List[Cart] = carts

  def create(items: List[CartEntry]): Cart = {
    val newCart = Cart(carts.size + 1L, items)
    carts = carts :+ newCart
    newCart
  }

  def findById(id: Long): Option[Cart] = {
    carts.find(_.id == id)
  }

  def delete(id: Long): Unit = {
    carts = carts.filterNot(_.id == id)
  }

  def update(id: Long, items: List[CartEntry]): Option[Cart] = {
    findById(id).map { _ =>
      val updatedCart = Cart(id, items)
      carts = carts.map(c => if (c.id == id) updatedCart else c)
      updatedCart
    }
  }
}
