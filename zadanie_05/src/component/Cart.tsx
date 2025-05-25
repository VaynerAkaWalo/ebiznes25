import {useCart} from "../context/CartContext.tsx";
import {Link} from "react-router-dom";

export function Cart() {
    const {entries, totalPrice} = useCart()

    return (
        <div className="h-1/2 w-1/2 flex flex-col border-2">
            <p className="text-center font-bold py-2 border-b-2">Koszyk</p>
            <ul className="flex py-2 border-b-2 [&>li]:py-2 [&>li]:px-10">
                <li className="w-2/5">Nazwa</li>
                <li>Ilosc</li>
                <li className="ml-auto">Cena</li>
            </ul>
            <ul className="flex flex-col">
                {entries.map(entry =>
                    <li key={entry.product.id} className="flex [&>div]:py-2 [&>div]:px-10">
                        <div className="w-2/5">{entry.product.Name}</div>
                        <div>{entry.quantity}</div>
                        <div className="ml-auto">{entry.product.Price * entry.quantity}zł</div>
                    </li>
                )}
            </ul>
            <div className="border-t-2 mt-auto py-2 px-4 flex justify-end items-center">
                <p>Razem:</p>
                <p className="w-1/5 text-right">{totalPrice()}zł</p>
            </div>
            {entries.length > 0 && (
                <div className="border-t-2 py-2 px-4 flex justify-center items-center">
                    <Link to="/payment">
                        <button className="button border-2 px-12 py-2">
                            Przejdź do płatności
                        </button>
                    </Link>
                </div>
            )}
        </div>
    )
}
