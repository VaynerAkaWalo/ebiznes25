import {Link} from "react-router-dom";

export function Header() {
    return (
        <ul className="flex flex-row w-full justify-around py-4 border-b-2">
            <li><Link to="/">Produkty</Link></li>
            <li><Link to="/cart">Koszyk</Link></li>
        </ul>
    )
}
