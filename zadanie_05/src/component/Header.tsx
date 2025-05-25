import {Link} from "react-router-dom";

export function Header() {
    return (
        <ul className="flex flex-row w-full justify-around py-4 border-b-2">
            <li><Link to="/" data-test-id="nav-link-products">Produkty</Link></li>
            <li><Link to="/cart" data-test-id="nav-link-cart">Koszyk</Link></li>
        </ul>
    )
}
