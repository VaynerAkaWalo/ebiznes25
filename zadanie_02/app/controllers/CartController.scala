package controllers

import models.{Cart, Category, ErrorMessage}
import play.api.libs.json._
import play.api.mvc._

import javax.inject._

@Singleton
class CartController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  def listAll(): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(Cart.all()))
  }

  def getById(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] =>
    Cart.findById(id) match {
      case Some(cart: Cart) => Ok(Json.toJson(cart))
      case None => NotFound(Json.toJson(ErrorMessage("Cart not found")))
    }
  }

  def delete(id: Long): Action[AnyContent] = Action { implicit request: Request[AnyContent] => {
    Cart.delete(id)
    NoContent
  }
  }

  def update(id: Long): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Cart].fold(
      errors => BadRequest(Json.toJson(ErrorMessage("Invalid message body " + errors))),
      cart => {
        Cart.update(id, cart.items) match {
          case Some(updatedCart) => Ok(Json.toJson(updatedCart))
          case None => NotFound(Json.toJson(ErrorMessage("Cart not found")))
        }
      }
    )
  }

  def create(): Action[JsValue] = Action(parse.json) { request =>
    request.body.validate[Cart].fold(
      errors => BadRequest(Json.toJson(ErrorMessage("Invalid message body " + errors))),
      cart => {
        val newCart = Cart.create(cart.items)
        Created(Json.toJson(newCart))
      }
    )
  }
}
