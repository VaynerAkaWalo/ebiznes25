package controllers

import models.{Category, ErrorMessage, Product}
import play.api.libs.json._
import play.api.mvc._

import javax.inject._

@Singleton
class CategoryController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  def listAll(): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(Category.all()))
  }

  def getById(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Category.findById(id) match {
      case Some(category: Category) => Ok(Json.toJson(category))
      case None => NotFound(Json.toJson(ErrorMessage("Category not found")))
    }
  }

  def delete(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] => {
    Category.delete(id)
    NoContent
  }
  }

  def update(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest(Json.toJson(ErrorMessage("Invalid message body " + errors))),
      category => {
        Category.update(id, category.name) match {
          case Some(updatedCategory) => Ok(Json.toJson(updatedCategory))
          case None => NotFound(Json.toJson(ErrorMessage("Category not found")))
        }
      }
    )
  }

  def create(): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Category].fold(
      errors => BadRequest(Json.toJson(ErrorMessage("Invalid message body " + errors))),
      category => {
        val newCategory = Category.create(category.name)
        Created(Json.toJson(newCategory))
      }
    )
  }
}
