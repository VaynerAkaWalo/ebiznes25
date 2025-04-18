package controllers

import models.{ErrorMessage, Product}
import play.api.libs.json._
import play.api.mvc._

import javax.inject._

@Singleton
class ProductController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  def listAll(): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(Product.all()))
  }

  def getById(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Product.findById(id) match {
      case Some(product: Product) => Ok(Json.toJson(product))
      case None => NotFound(Json.toJson(ErrorMessage("Product not found")))
    }
  }

  def delete(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] => {
    Product.delete(id)
    NoContent
  }
  }

  def update(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest(Json.toJson(ErrorMessage("Invalid message body " + errors))),
      product => {
        Product.update(id, product.name, product.price) match {
          case Some(updatedProduct) => Ok(Json.toJson(updatedProduct))
          case None => NotFound(Json.toJson(ErrorMessage("Product not found")))
        }
      }
    )
  }

  def create(): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Product].fold(
      errors => BadRequest(Json.toJson(ErrorMessage("Invalid message body " + errors))),
      product => {
        val newProduct = Product.create(product.name, product.price)
        Created(Json.toJson(newProduct))
      }
    )
  }
}
