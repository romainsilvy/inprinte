import React from 'react';
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import bar2 from '../images/bar2.svg'
import random from '../images/random.jpg'
import bar_panier from "../images/bar_panier.svg"
import add_panier from "../images/add_panier.svg"
import loading from "../images/loading.svg"

import "../styles/Panier.css"

class Panier extends React.Component {

    constructor(props) {
        super(props);

        this.state = {
            items: [],
            DataisLoaded: false
        };
    }

    componentDidMount() {
        let cart = localStorage.getItem("cart");
        let emptyCart = false
        if (cart) {
            emptyCart = false
        } else {
            emptyCart = true;
        }
        if (!emptyCart) {
            fetch("http://localhost:8080/panier?cart=" + cart, {
                method: 'GET'})
                .then((res) => res.json())
                .then((json) => {
                    this.setState({
                        items: json,
                        DataisLoaded: true
                    });
                })
        }
    }


    render() {
        const { DataisLoaded, items } = this.state;
        if (!DataisLoaded) return (
            <>
                <Navbar />
                <img className="loading" src={loading} alt="loading" />
            </>
        )
        return (
        <>
            <Navbar />
            <div className="panier">
                <div className="titleCat1Cart">
                    <h1>Panier</h1>
                    <img className="bar2" src={bar2} alt="bar" />
                </div>
                <div className="allCartProducts">
                    {
                            items.allItems.map((item) => (
                                <a href={"produit/" + item.id} className="article">
                                    <img className="picture" src={item.picture} alt="image" />
                                    <div className="infoArticle">
                                        <h2>{item.name}</h2>
                                        <p>{item.description}</p>
                                    </div>
                                    <div className="status">
                                        <p className="color">Quantité {item.quantity}</p>
                                        <p>{item.price * item.quantity}€</p>
                                    </div>
                                </a>
                            ))
                        }
                </div>

                {/* <div className="article2">
                    <img className="picture" src={random} alt="image" />
                    <div className="infoArticle">
                        <h2>Nom du produit</h2>
                        <p>Description</p>
                    </div>
                    <div className="status">
                        <p className="color">Quantité 1</p>
                        <p>135€</p>
                    </div>
                </div> */}
            </div>
            <div className="total">
                <img className="bar_panier" src={bar_panier} alt="image" />
                <div className="total2">
                    <p className="infoTotal">Total </p>
                    <p className="infoTotal">{items.totalPrice}€</p>
                </div>
                <img className="add_panier" src={add_panier} alt="image" />
            </div>
            <Footer />
        </>
    );
    }
};

export default Panier;