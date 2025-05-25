import {useState, MouseEvent} from "react";
import {CardPayment} from "./CardPayment.tsx";
import {BackendClient} from "../client/BackendClient.ts";

interface PaymentProvider {
  name: string
  type: PaymentProviderType
}

interface CardDetails {
  firstname: string
  lastname: string
  cardNumber: number
  expireTime: string
  cvv: number
}

const initialCardDetails: CardDetails = {
  firstname: "",
  lastname: "",
  cardNumber: 0,
  expireTime: "",
  cvv: 0
}

enum PaymentProviderType {
  CARD_PAYMENT,
}

export function PaymentPage() {
  const [selectedProvider, setSelectedProvider] = useState<number>(-1)
  const [cardDetails, setCardDetails] = useState<CardDetails>(initialCardDetails)
  const [validationError, setValidationError] = useState<string>("")

  const providers: PaymentProvider[] = [
    {
      name: 'Visa',
      type: PaymentProviderType.CARD_PAYMENT
    },
    {
      name: 'MasterCard',
      type: PaymentProviderType.CARD_PAYMENT
    }
  ]

  const handleSelectProvider = (e: MouseEvent<HTMLButtonElement>) => {
    setSelectedProvider(Number(e.currentTarget.value))
    setValidationError("")
  }

  const handlePaymentSuccess = () => {
    setSelectedProvider(-1)
    setCardDetails(initialCardDetails)
    setValidationError("")
  }

  const handleCardDetailsChange = (name: string, value: string) => {
    setCardDetails(prev => ({
      ...prev,
      [name]: value
    }))
    setValidationError("")
  }

  const validateCardDetails = (): boolean => {
    if (!cardDetails.firstname) {
      setValidationError("Podaj imie")
      return false
    }
    if (!cardDetails.lastname) {
      setValidationError("Podaj nazwisko")
      return false
    }
    if (cardDetails.cardNumber.toString().length !== 16) {
      setValidationError("Nieprawidłowy numer karty")
      return false
    }
    if (cardDetails.cvv.toString().length !== 3) {
      setValidationError("Nieprawidłowy kod CVV")
      return false
    }
    if (cardDetails.expireTime.length !== 5) {
      setValidationError("Nieprawidłowy format daty ważności")
      return false
    }

    return true
  }

  const handlePay = () => {
    if (!validateCardDetails()) {
      return
    }

    const req = {
      providerId: providers[selectedProvider].name,
      amount: 0,
      card: {
        firstName: cardDetails.firstname,
        lastName: cardDetails.lastname,
        cardNumber: Number(cardDetails.cardNumber),
        expireTime: cardDetails.expireTime,
        CVV: Number(cardDetails.cvv)
      }
    }

    BackendClient.payByCard(req)
      .then(() => {
        handlePaymentSuccess()
      })
      .catch(error => {
        console.error("Payment failed:", error)
        setValidationError("Bład połączenia z serwerem")
      });
  }

  const paymentProviderList = () => {
    return (
      <div className="flex flex-col w-full h-full text-center">
        <p className="w-full py-3 border-b-2">Select payment method</p>
        <ul className="flex flex-col w-full h-full justify-evenly">
          {providers.map((provider, index) =>
            <li className="flex mx-auto h-1/5 w-4/5 border-2 justify-center items-center button">
              <button value={index} className="flex h-full w-full justify-center items-center" onClick={handleSelectProvider}>{provider.name}</button>
            </li>
          )}
        </ul>
      </div>
    )
  }

  const paymentModal = () => {
    const provider: PaymentProvider = providers[selectedProvider]

    return (
      <div className="flex flex-col w-full h-full">
        <p className="w-full py-3 text-center border-b-2">{provider.name}</p>
        <div className="flex w-full h-full">
          <CardPayment
            cardDetails={cardDetails}
            onCardDetailsChange={handleCardDetailsChange}
          />
        </div>
        {validationError && (
          <div className="flex w-full">
            <p className="text-red-500 text-sm text-center w-full">{validationError}</p>
          </div>
        )}
        <div className="flex w-full py-4 justify-around [&>button]:border-2 [&>button]:w-1/3">
          <button className="button" onClick={() => setSelectedProvider(-1)}>Cofnij</button>
          <button className="button" onClick={handlePay}>Zapłać</button>
        </div>
      </div>
    )
  }

  return (
    <div data-test-id="payment-page" className="flex w-1/4 h-1/2 border-2">
      {selectedProvider === -1 && paymentProviderList()}
      {selectedProvider !== -1 && paymentModal()}
    </div>
  )
}
