package models

import play.api.libs.json.{Format, Json}

case class Product(id: Long, name: String, price: Double)

object Product {
  implicit val productFormat: Format[Product] = Json.format[Product]

  var products: List[Product] = List(
    Product(1, "Milk", 2.00),
    Product(2, "Candy", 5.00)
  )

  def all(): List[Product] = products

  def create(name: String, price: Double): Product = {
    val newProduct = Product(products.size + 1L, name, price)
    products = products :+ newProduct
    newProduct
  }

  def findById(id: Long): Option[Product] = {
    products.find(_.id == id)
  }

  def delete(id: Long): Unit = {
    products = products.filterNot(_.id == id)
  }

  def update(id: Long, name: String, price: Double): Option[Product] = {
    findById(id).map { _ =>
      val updatedTask = Product(id, name, price)
      products = products.map(p => if (p.id == id) updatedTask else p)
      updatedTask
    }
  }
}
