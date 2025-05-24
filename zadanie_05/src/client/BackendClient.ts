import {AxiosResponse} from "axios";
import {HTTPClient} from "./Client.ts";

export interface Product {
  Name: string
  id: string
  Price: number
}

export interface PaymentStatus {
  status: string
}

export interface CardPayment {
  providerId: string,
  amount: number,
  card: Card
}

export interface Card {
  firstName: string
  lastName: string
  cardNumber: number
  expireTime: string
  CVV: number
}


class Backend {
  private readonly baseUrl: string;

  constructor() {
    this.baseUrl = "http://localhost:8000";
  }

  public getProducts = async (): Promise<AxiosResponse<Product[]>> => {
    return HTTPClient.get(this.baseUrl + "/products")
  }

  public payByCard = async (req: CardPayment): Promise<AxiosResponse<PaymentStatus>> => {
    return HTTPClient.post(this.baseUrl + "/payment/card", req)
  }
}

export const BackendClient = new Backend();
