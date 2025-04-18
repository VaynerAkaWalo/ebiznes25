package models

import play.api.libs.json.{Format, Json}

case class Product(id: Long, name: String, price: Double, categoryId: Long)

object Product {
  implicit val productFormat: Format[Product] = Json.format[Product]

  var products: List[Product] = List(
    Product(1, "Milk", 2.00, 2),
    Product(2, "Candy", 5.00, 1)
  )

  def all(): List[Product] = products

  def create(name: String, price: Double, categoryId: Long): Product = {
    val newProduct = Product(products.size + 1L, name, price, categoryId)
    products = products :+ newProduct
    newProduct
  }

  def findById(id: Long): Option[Product] = {
    products.find(_.id == id)
  }

  def delete(id: Long): Unit = {
    products = products.filterNot(_.id == id)
  }

  def update(id: Long, name: String, price: Double, categoryId: Long): Option[Product] = {
    findById(id).map { _ =>
      val updatedTask = Product(id, name, price, categoryId)
      products = products.map(p => if (p.id == id) updatedTask else p)
      updatedTask
    }
  }
}
