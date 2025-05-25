import {ChangeEvent} from "react";

interface CardDetails {
  firstname: string
  lastname: string
  cardNumber: number
  expireTime: string
  cvv: number
}

interface CardPaymentProps {
  cardDetails: CardDetails
  onCardDetailsChange: (name: string, value: string) => void
}

export function CardPayment({ cardDetails, onCardDetailsChange }: CardPaymentProps) {
  const updateFirstName = (e: ChangeEvent<HTMLInputElement>) => {
    onCardDetailsChange('firstname', e.target.value)
  }

  const updateLastName = (e: ChangeEvent<HTMLInputElement>) => {
    onCardDetailsChange('lastname', e.target.value)
  }

  const updateCardNumber = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value
    if (!isNaN(Number(value)) && value.length <= 16) {
      onCardDetailsChange('cardNumber', value)
    }
  }

  const updateExpireTime = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value
    if (value.length < 2) {
      onCardDetailsChange('expireTime', value)
      return
    }

    if (value.length == 2) {
      onCardDetailsChange('expireTime', value + '/')
      return
    }

    if (value.length == 3) {
      onCardDetailsChange('expireTime', value)
      return
    }

    const [month, year] = e.target.value.split('/')
    if (Number(month) > 0 && Number(month) < 13 && (year.length == 1 || Number(year) > 25 && Number(year) < 100)) {
      onCardDetailsChange('expireTime', value)
    }
  }

  const updateCVV = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    if (!isNaN(Number(value)) && value.length < 4) {
      onCardDetailsChange('cvv', value)
    }
  }

  return (
    <div className="flex flex-col w-full items-center py-8 justify-around">
      <div className="flex">
        <div className="flex flex-col w-1/2 items-center">
          <p>Imię</p>
          <input className="border-1 w-2/3" value={cardDetails.firstname} onChange={updateFirstName} data-test-id="card-input-firstname"/>
        </div>
        <div className="flex flex-col w-1/2 items-center">
          <p>Nazwisko</p>
          <input className="border-1 w-2/3" value={cardDetails.lastname} onChange={updateLastName} data-test-id="card-input-lastname"/>
        </div>
      </div>
      <div className="flex flex-col py-4 w-full">
        <p className="text-center">Podaj numer karty</p>
        <div className="flex justify-around">
          <input className="border-1 w-1/2 text-center" placeholder="xxxx-xxxx-xxxx-xxxx" value={cardDetails.cardNumber} onChange={updateCardNumber} data-test-id="card-input-number"/>
        </div>
      </div>
      <div className="flex">
        <div className="flex flex-col w-1/2 items-center">
          <p>Data ważności</p>
          <input className="border-1 w-2/3 text-center" placeholder="MM/YY" value={cardDetails.expireTime} onChange={updateExpireTime} data-test-id="card-input-expire"/>
        </div>
        <div className="flex flex-col w-1/2 items-center">
          <p>Kod CVV</p>
          <input className="border-1 w-2/3 text-center" placeholder="XXX" value={cardDetails.cvv} onChange={updateCVV} data-test-id="card-input-cvv"/>
        </div>
      </div>
    </div>
  )
}
