import React, { useState, useEffect } from 'react'
import { Navbar } from '../components/Navbar';
import { Footer } from '../components/Footer';
import loading from '../images/loading.svg';

import { ShopProduct } from '../components/ShopProduct';
import "../styles/Produit.css"
import random from '../images/random.jpg'
import btn_add from '../images/btn_add.svg'
import like from '../images/like.svg'
import bar2 from '../images/bar2.svg'
import {useParams} from 'react-router-dom';

function Produit() {
    const { id } = useParams();
    // console.log(id)
    const [state, setState] = useState({
        items: [],
        DataisLoaded: false
    })

    useEffect(() => {     
        fetch("http://localhost:8080/produit/"+id)
            .then((res) => res.json())
            .then((json) => {
                setState({
                    items: json,
                    DataisLoaded: true
                });
            })

    }, [state.DataisLoaded]);
    if (!state.DataisLoaded) return (
        <>
            <Navbar />
            <img className="P-loading" src={loading} alt="loading" />
        </>
    )
    return (
        <>
            <Navbar />
            <div className="P-containerProduit">
                <img className="P-produitPicture" src={state.items.product_data.picture} alt="image" />
                <div className="P-infosProduit">
                    <h1>{state.items.product_data.name}</h1>
                    <img className="P-bar2" src={bar2} alt="bar" />
                    <p>{state.items.product_data.description}</p>
                    <p>{state.items.product_data.price}€</p>
                    <img className="P-btn_add" src={btn_add} alt="btn" />
                    <img className="P-like" src={like} alt="btn" />
                </div>
            </div>
            <div className="P-containerProduit2">
                <div className="P-produit">
                    <div className="P-titleProduit">
                        <h1>D'autres ont également achetés</h1>
                        <img className="P-bar2" src={bar2} alt="bar" />
                    </div>
                </div>
                <div className="P-allShopProducts">
                    {
                        state.items.related.map((item) => (
                            <ShopProduct product={item} />
                        ))
                    }
                </div>
            </div>
            <Footer />
        </>
    )
}
export default Produit;