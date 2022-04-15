import React from 'react';
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import bar2 from '../images/bar2.svg'
import random from '../images/random.jpg'
import bar_panier from "../images/bar_panier.svg"
import add_panier from "../images/add_panier.svg"

import "../styles/Panier.css"
const panier = () => {
    return (
        <>
            <Navbar />
            <div className="panier">
                <div className="titleCat1">
                    <h1>Panier</h1>
                    <img className="bar2" src={bar2} alt="bar" />
                </div>
                <div className="article1">
                    <img className="picture" src={random} alt="image" />
                    <div className="infoArticle">
                        <h2>Nom du produit</h2>
                        <p>Description</p>
                    </div>
                    <div className="status">
                        <p className="color">Quantité 1</p>
                        <p>135€</p>
                    </div>
                </div>
                <div className="article2">
                    <img className="picture" src={random} alt="image" />
                    <div className="infoArticle">
                        <h2>Nom du produit</h2>
                        <p>Description</p>
                    </div>
                    <div className="status">
                        <p className="color">Quantité 1</p>
                        <p>135€</p>
                    </div>
                </div>
            </div>
            <div className="total">
                <img className="bar_panier" src={bar_panier} alt="image" />
                <div className="total2">
                    <p className="infoTotal">Total </p>
                    <p className="infoTotal">100€</p>
                </div>
                <img className="add_panier" src={add_panier} alt="image" />
            </div>
            <Footer />
        </>
    );
};

export default panier;