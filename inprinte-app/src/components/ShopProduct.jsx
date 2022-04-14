import {appendScript} from '../utils/appendScript'
import {addToCart, clearCart} from '../scripts/test'

import React from 'react';


export class ShopProduct extends React.Component {

    constructor(props) {
        super(props);
    }

    componentDidMount () {
        clearCart()
    }

    render() {

    console.log(this.props)
    let url = "/produit/" + this.props.product.Id_product
    return (
        <>
             <a className="shopProduct">
                <img className="productPicture" src="https://picsum.photos/400/400" alt="product image" />
                <div className="productInfos">
                    <p className="productName">{this.props.product.Name}</p>
                    <p className="productPrice">{this.props.product.Price} €</p>
                </div>
                <p className="productDescription">
                    {this.props.product.Description}
                </p>
                <p className="addToCart">add to cart</p>
            </a>
        </>
    );
    }
}


export default ShopProduct;