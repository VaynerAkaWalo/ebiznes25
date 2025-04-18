package models

import play.api.libs.json.{Format, Json}

case class Category(id: Long, name: String)

object Category {
  implicit val categoryFormat: Format[Category] = Json.format[Category]

  var categories: List[Category] = List(
    Category(1, "Sweets"),
    Category(2, "Dairy")
  )

  def all(): List[Category] = categories

  def create(name: String): Category = {
    val newCategory = Category(categories.size + 1L, name)
    categories = categories :+ newCategory
    newCategory
  }

  def findById(id: Long): Option[Category] = {
    categories.find(_.id == id)
  }

  def delete(id: Long): Unit = {
    categories = categories.filterNot(_.id == id)
  }

  def update(id: Long, name: String): Option[Category] = {
    findById(id).map { _ =>
      val updatedCategory = Category(id, name)
      categories = categories.map(c => if (c.id == id) updatedCategory else c)
      updatedCategory
    }
  }
}
